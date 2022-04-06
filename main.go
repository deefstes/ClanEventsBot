package main

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/kenshaw/baseconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	buildNumber     string
	liveTime        time.Time
	config          Configuration
	mongoClient     *mongo.Client
	discordSession  *discordgo.Session
	defaultLocation *time.Location
	guildVars       map[string]*GuildVars
	ErrNoRecords    = errors.New("no records")
)

type GuildVars struct {
	guild        Guild
	impersonated ClanUser
	timezones    []TimeZone
	tzByAbbr     map[string]TimeZone
	tzByEmoji    map[string]TimeZone
	escrowEvents map[string]DevelopingEvent
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "%s - FATAL ERROR: %+v\r\n", time.Now().Format("2006-01-02 15:04:05"), r)
		}
	}()

	if buildNumber == "" {
		buildNumber = "N/A"
	}

	liveTime = time.Now()
	guildVars = make(map[string]*GuildVars)
	var err error

	// Read config file
	config, err = ReadConfig()
	if err != nil {
		fmt.Println("Error reading config file")
		return
	}

	defaultLocation, _ = time.LoadLocation("Europe/London")

	// Connect to MongoDB
	mongoClient, err = mongo.NewClient(options.Client().ApplyURI(config.MongoDB))
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	err = mongoClient.Connect(ctx)
	if err != nil {
		panic(fmt.Errorf("connecting to mongodb: %w", err))
	}
	defer mongoClient.Disconnect(ctx)

	// Attempt connecting to the database 4 times to allow for network resources not yet being available due to system startup
	interval := 15
	for i := 0; i < 4; i++ {
		fmt.Println("attempting db connection")
		// ctx, c := context.WithTimeout(context.Background(), 10*time.Second)
		// defer c()
		err = mongoClient.Ping(ctx, readpref.Primary())
		if err == nil {
			break
		} else {
			fmt.Printf("db connection failed, waiting %d seconds\r\n", interval)
			time.Sleep(time.Duration(interval) * time.Second)
		}
		interval = interval * 2 // Make interval between successive attempts longer
	}
	if err != nil {
		panic(fmt.Errorf("unable to ping mongodb: %w", err))
	}

	var guilds []Guild
	collection := mongoClient.Database("ClanEvents").Collection("Guilds")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		fmt.Printf("Error reading guilds: %v\r\n", err)
	}
	if err = cur.All(context.TODO(), &guilds); err != nil {
		fmt.Printf("Error decoding guilds: %v\r\n", err)
	}
	if len(guilds) == 0 {
		fmt.Printf("No registered guilds\r\n")
	}

	for _, guild := range guilds {
		guildVars[guild.ID] = NewGuildVars(guild)
	}

	// Create a new Discord session using the provided bot token.
	discordSession, err = discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	defer discordSession.Close()

	discordSession.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuilds | discordgo.IntentsGuildMembers | discordgo.IntentsGuildMessages | discordgo.IntentsGuildMessageReactions)

	// for _, guild := range guilds {
	// 	members, _ := discordSession.GuildMembers(guild.ID, "", 1000)
	// 	for _, member := range members {
	// 		m, _ := discordSession.GuildMember(guild.ID, member.User.ID)
	// 		fmt.Printf("%s: %s\r\n", guild.Name, m.User.Username)
	// 	}
	// }

	// Set up a ticker that triggers a service routine every n minutes
	ticker := time.NewTicker(time.Minute * config.ServiceTimer)
	defer ticker.Stop()
	go func() {
		for {
			<-ticker.C
			serviceRoutine()
		}
	}()

	// Register the messageCreate and messageReact functions as callbacks for MessageCreate and MessageReactionAdd events.
	discordSession.AddHandler(messageCreate)
	discordSession.AddHandler(messageReact)

	// Open a websocket connection to Discord and begin listening.
	err = discordSession.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Printf("ClanEventsBot (build number %s) is now running", buildNumber)
	if config.DebugLevel > 0 {
		fmt.Printf(" with DebugLevel=%d\r\n", config.DebugLevel)
	} else {
		fmt.Println()
	}
	fmt.Println("Press CTRL-C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func NewGuildVars(g Guild) *GuildVars {
	ee := make(map[string]DevelopingEvent)

	var tzs []TimeZone
	ctz := mongoClient.Database(fmt.Sprintf("ClanEvents%s", g.ID)).Collection("TimeZones")

	sortopts := options.Find().SetSort(bson.D{{"abbrev", 1}})
	cur, err := ctz.Find(context.Background(), bson.D{}, sortopts)
	if err != nil {
		fmt.Printf("Error reading timezones: %v\r\n", err)
		return nil
	}
	if err = cur.All(context.TODO(), &tzs); err != nil {
		fmt.Printf("Error reading timezones: %v\r\n", err)
		return nil
	}
	tzBA, tzBE := constructTZMaps(tzs)

	return &GuildVars{
		guild:        g,
		timezones:    tzs,
		tzByAbbr:     tzBA,
		tzByEmoji:    tzBE,
		escrowEvents: ee,
	}
}

func constructTZMaps(tzs []TimeZone) (tzBA map[string]TimeZone, tzBE map[string]TimeZone) {
	tzBA = make(map[string]TimeZone)
	tzBE = make(map[string]TimeZone)

	for _, timezone := range tzs {
		tzBA[timezone.Abbrev] = timezone
		if timezone.Emoji != "" {
			bytearray, err := hex.DecodeString(timezone.Emoji)
			if err != nil {
				panic(err)
			}
			emojistr := string(bytearray[:len(bytearray)])
			tzBE[emojistr] = timezone
		}
	}

	return tzBA, tzBE
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "%+v", r)
			message := fmt.Sprintf("Well this is embarrasing :flushed:.")
			message = fmt.Sprintf("%s\r\nSomething went wrong and I don't know what it is. We shall never speak of this again.", message)
			sendMessage(m.ChannelID, message)
		}
	}()

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Ignore any messages created by other bots
	if m.Author.Bot {
		return
	}

	// If messages is not a command, bail out
	if !strings.HasPrefix(m.Content, config.CommandPrefix) {
		return
	}
	command := strings.TrimPrefix(m.Content, config.CommandPrefix)
	commandElements := getArgs(command)

	// If message does not contain guild information, bail out
	guild := getGuild(s, m)
	if guild == nil {
		return
	}

	if config.DebugLevel > 0 {
		fmt.Printf("Guild=%s, Author=%s(%s), Command=%s\r\n", guild.Name, m.Author.Username, m.Author.ID, command)
	}

	if strings.HasPrefix(command, "help") {
		BotHelp(guild, s, m, commandElements)
	} else if strings.HasPrefix(command, "listevents") {
		ListEvents(guild, s, m, commandElements)
	} else if strings.HasPrefix(command, "newevent ") {
		NewEvent(guild, s, m, commandElements)
	} else if strings.HasPrefix(command, "new ") {
		New(guild, s, m, commandElements)
	} else if strings.HasPrefix(command, "edit ") {
		Edit(guild, s, m, commandElements)
	} else if strings.HasPrefix(command, "cancel ") {
		CancelEvent(guild, s, m, commandElements)
	} else if strings.HasPrefix(command, "cancelevent ") {
		CancelEvent(guild, s, m, commandElements)
	} else if strings.HasPrefix(command, "signup ") {
		Signup(guild, s, m, commandElements)
	} else if strings.HasPrefix(command, "leave ") {
		Leave(guild, s, m, commandElements)
	} else if strings.HasPrefix(command, "impersonate ") {
		Impersonate(guild, s, m, commandElements)
	} else if strings.HasPrefix(command, "unimpersonate") {
		Unimpersonate(guild, s, m, commandElements)
	} else if strings.HasPrefix(command, "details ") {
		Details(guild, s, m, commandElements)
	} else if strings.HasPrefix(command, "test") {
		Test(guild, s, m, commandElements)
	} else if strings.HasPrefix(command, "say") || strings.HasPrefix(command, "echo") {
		Echo(guild, s, m, commandElements)
	} else if strings.HasPrefix(command, "wisdom") {
		Wisdom(guild, s, m, commandElements)
	} else if strings.HasPrefix(command, "addnaughtylist ") {
		AddNaughty(guild, s, m, commandElements)
	} else if strings.HasPrefix(command, "removenaughtylist ") {
		RemoveNaughty(guild, s, m, commandElements)
	} else if strings.HasPrefix(command, "naughtylist") {
		ListNaughty(guild, s, m, commandElements)
	} else if strings.HasPrefix(command, "addserver") {
		AddServer(guild, s, m, commandElements)
	} else if strings.HasPrefix(command, "addtimezone ") {
		AddTimeZone(guild, s, m, commandElements)
	} else if strings.HasPrefix(command, "removetimezone ") {
		RemoveTimeZone(guild, s, m, commandElements)
	} else if strings.HasPrefix(command, "listtimezones") {
		ListTimeZones(guild, s, m, commandElements)
	} else if strings.HasPrefix(command, "roletimezone ") {
		RoleTimeZone(guild, s, m, commandElements)
	} else {
		message := fmt.Sprintf("If word gets out, everyone will want an extra pancreas so I think they must... uhm, sorry, what? Were you talking to me? Can't you see I'm busy? Anyway, I don't know what **%s** means. It's certainly not a command that I've been programmed with.", commandElements[0])
		message = fmt.Sprintf("%s\r\nFor a list of valid commands, type the following:\r\n```%shelp```", message, config.CommandPrefix)
		sendMessage(m.ChannelID, message)
	}
}

// This function will be called (due to AddHandler above) every time a new
// reaction is added to a message on any channel that the autenticated bot has access to.
func messageReact(s *discordgo.Session, m *discordgo.MessageReactionAdd) {

	// Ignore all reactions added by the bot itself
	if m.UserID == s.State.User.ID {
		return
	}

	ProcessReaction(s, m)
}

func sendMessage(channelID string, message string) {
	discordSession.ChannelMessageSend(channelID, message)
}

func serviceRoutine() {
	c := mongoClient.Database("ClanEvents").Collection("Guilds")

	var guilds []Guild

	cur, err := c.Find(context.Background(), bson.D{})
	if err != nil {
		fmt.Printf("Error reading guilds: %v\r\n", err)
		return
	}
	if err = cur.All(context.TODO(), &guilds); err != nil {
		fmt.Printf("Error decoding guilds: %v\r\n", err)
		return
	}
	if len(guilds) == 0 {
		fmt.Printf("No registered guilds\r\n")
		return
	}

	for _, guild := range guilds {
		serviceGuild(guild)
	}
}

func serviceGuild(guild Guild) {
	deliverInsult(guild.ID)
	archiveEvents(guild.ID)
}

func getArgs(s string) []string {
	re := regexp.MustCompile("\".+?\"|“.+?”|'.+?'|\\S+")
	args := re.FindAllString(s, -1)
	for i := 0; i < len(args); i++ {
		args[i] = strings.TrimPrefix(args[i], "\"")
		args[i] = strings.TrimSuffix(args[i], "\"")
		args[i] = strings.TrimPrefix(args[i], "“")
		args[i] = strings.TrimSuffix(args[i], "”")
		args[i] = strings.TrimPrefix(args[i], "'")
		args[i] = strings.TrimSuffix(args[i], "'")
	}
	return args
}

func archiveEvents(guildID string) {
	filter := bson.M{}
	filter["dateTime"] = bson.M{
		"$lte": time.Now().Add(-1 * time.Hour),
	}
	filter["archived"] = bson.M{
		"$ne": true,
	}

	c := mongoClient.Database(fmt.Sprintf("ClanEvents%s", guildID)).Collection("Events")

	var results []ClanEvent
	sortopts := options.Find().SetSort(bson.D{{"dateTime", 1}})
	cur, err := c.Find(context.Background(), filter, sortopts)
	if err != nil {
		fmt.Printf("Error reading unarchived events on guild %s: %v\r\n", guildID, err)
		return
	}
	if err = cur.All(context.TODO(), &results); err != nil {
		fmt.Printf("Error decoding unarchived events on guild %s: %v\r\n", guildID, err)
		return
	}

	for _, event := range results {
		upsertfilter := bson.M{"eventId": event.EventID}
		event.Archived = true
		event.EventID = fmt.Sprintf("%s_%s", time.Now().Format("060102150405"), event.EventID)
		event.ObjectID = primitive.NilObjectID
		_, err := c.ReplaceOne(
			context.Background(),
			upsertfilter,
			event,
			options.Replace().SetUpsert(true),
		)
		if err != nil {
			fmt.Printf("Error archiving event %s on guild %s: %v\r\n", event.EventID, guildID, err)
			return
		}
	}
}

func deliverInsult(guildID string) {
	// Find default channel and insultee in DB
	c := mongoClient.Database(fmt.Sprintf("ClanEvents%s", guildID)).Collection("Config")
	var config ClanConfig
	rslt := c.FindOne(context.Background(), bson.D{})
	if rslt.Err() == mongo.ErrNoDocuments {
		return
	}
	if rslt.Err() != nil {
		fmt.Printf("Error reading config for guild %s: %v", guildID, rslt.Err())
		return
	}
	err := rslt.Decode(&config)
	if err != nil {
		fmt.Printf("Error decoding config for guild %s: %v", guildID, err)
		return
	}

	prob := rand.Float64()
	if prob <= config.InsultProbability {
		insultee, err := getInsultee(guildID)
		if err == ErrNoRecords {
			return
		} else if err != nil {
			fmt.Printf("Error delivering insult on guild %s\r\n", guildID)
			return
		}
		message := getInsult(insultee.Mention())
		if message != "" {
			sendMessage(config.DefaultChannel, message)
		}
	}
}

func getInsultee(guildID string) (ClanUser, error) {
	c := mongoClient.Database(fmt.Sprintf("ClanEvents%s", guildID)).Collection("NaughtyList")

	var insultees []ClanUser
	var insultee ClanUser

	cur, err := c.Find(context.Background(), bson.D{})
	if err != nil {
		return insultee, fmt.Errorf("reading insultees: %w", err)
	}
	if err = cur.All(context.TODO(), &insultees); err != nil {
		return insultee, fmt.Errorf("decoding insultees: %w", err)
	}
	if len(insultees) == 0 {
		return insultee, ErrNoRecords
	}

	index := rand.Intn(len(insultees))
	insultee = insultees[index]

	return insultee, nil
}

func getCanonicalInsult(mention string) string {
	canon := []string{
		fmt.Sprintf("%s is a stuck up, half-witted, scruffy-looking… Nerf herder!", mention),
		fmt.Sprintf("%s's parents are living proof that two wrongs don't make a right.", mention),
		fmt.Sprintf("EventsBot farts in %s's general direction.", mention),
		fmt.Sprintf("%s's brain is so minute that if a hungry cannibal cracked that head open, there wouldn't be enough to cover a small water biscuit.", mention),
		fmt.Sprintf("%s was fired as a bank clerk when a lady asked to check her balance and %s pushed her over.", mention, mention),
		fmt.Sprintf("It takes %s two minutes to cook minute rice", mention),
		fmt.Sprintf("%s's momma is so fat, she was baptised at SeaWorld.", mention),
		fmt.Sprintf("%s is so ugly that onions cry at the sight.", mention),
		fmt.Sprintf("%s has so little sex appeal, even hookers claim to have a headache.", mention),
		fmt.Sprintf("When %s was born the doctor saw the face, then saw the arse, and exclaimed \"Look... twins!\"", mention),
		fmt.Sprintf("%s was born on the highway, which is where most accidents happen.", mention),
		fmt.Sprintf("%s tiptoes past the medicine cabinet so as not to wake the sleeping pills.", mention),
		fmt.Sprintf("The village called, they want their idiot back. Has anyone seen %s?", mention),
		fmt.Sprintf("%s's family tree could better be described as a cactus. They're all a bunch of pricks.", mention),
		fmt.Sprintf("%s and two mates walked into a bar. You'd think at least one of them would've seen it.", mention),
		fmt.Sprintf("\"Honest m'lud\", said %s to the judge, \"I was just helping the sheep over the fence.\"", mention),
		fmt.Sprintf("%s licks the bus' windows.", mention),
		fmt.Sprintf("Family Law Quiz:\r\n%s is married without a prenuptual agreement.\r\nHe has 3 children from a previous marriage.\r\nShe has 4 children from two previous marriages and another two of which the father is unknown.\r\nThey own a trailer which her father gifted to them but the contents was assembled from various things he's stolen over the years.\r\nThey were both underaged when they married.\r\nThey now file for divorce.\r\nYes, or no; Are they still cousins?", mention),
		fmt.Sprintf("%s thinks a seven-course meal is a badger roadkill and a six-pack.", mention),
		fmt.Sprintf("%s is not the brightest candle in the chandelier.", mention),
		fmt.Sprintf("%s is overdrawn at the memory bank.", mention),
		fmt.Sprintf("If %s's brains were ink, you couldn't dot an i.", mention),
		fmt.Sprintf("%s couldn't ride a nightmare without falling out of bed.", mention),
		fmt.Sprintf("They had to burn down the school to get %s out of fifth grade.", mention),
		fmt.Sprintf("If laughter is the best medicine, %s's face must be curing the world.", mention),
		fmt.Sprintf("With that hideous face, %s scares the crap out of the toilet.", mention),
		fmt.Sprintf("The only way %s will ever get laid is by crawling up a chicken's ass and waiting it out.", mention),
		fmt.Sprintf("EventsBot is jealous of people who don't know %s.", mention),
		fmt.Sprintf("Just remember folks, brains aren't everything. In %s's case, they're nothing.", mention),
		fmt.Sprintf("Look, I understand that some babies were dropped on their heads. But seriously, %s was clearly thrown at a wall.", mention),
		fmt.Sprintf("Hey %s, why don't you slip into something more comfortable... like a coma.", mention),
		fmt.Sprintf("%s is the reason the gene pool needs a lifeguard.", mention),
		fmt.Sprintf("I've seen people like %s before, but I had to pay an admission.", mention),
		fmt.Sprintf("Turns out that sex position can affect the intelligence of the conceived baby... and clearly %s's parents used the wrong one.", mention),
	}

	index := rand.Intn(len(canon))
	return canon[index]
}

func getEvilInsult(mention string) string {
	url := "https://evilinsult.com/generate_insult.php?lang=en&type=text"

	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return ""
		}
		bodyString := string(bodyBytes)
		return fmt.Sprintf("%s, %s", mention, bodyString)
	}

	return ""
}

func getMattbasInsult(mention string) string {
	url := "https://insult.mattbas.org/api/insult.txt?template=is+as+<adjective>+as+<article+target=adj1>+<adjective+min=1+max=1+id=adj1>+<amount>+of+<adjective+min=1+max=1>+<animal>+<animal_part>"

	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return ""
		}
		bodyString := string(bodyBytes)
		return fmt.Sprintf("%s %s", mention, bodyString)
	}

	return ""
}

func getInsult(mention string) string {
	source := rand.Intn(3)
	switch source {
	case 0: // Insult from local canon
		return getCanonicalInsult(mention)
	case 1: // Insult from https://evilinsult.com/api/?ref=public-apis
		return getEvilInsult(mention)
	case 2: // Insult from https://insult.mattbas.org/api/
		return getMattbasInsult(mention)
	}

	return ""
}

func getEventID(t time.Time) string {
	timevalue := (t.Day() + t.Hour()*32 + t.Minute()*32*24 + t.Second()*32*24*60)
	timevalue = mapInt(timevalue, 2764799, 1679615) // 2764799=max number that timevalue can be (d=31, h=23, m=59, s=59), 1679615=max value that we want (encodes to ZZZZ) to ensure 4 digit ID
	newid, _ := baseconv.Encode36FromDec(fmt.Sprintf("%d", timevalue))
	newid = strings.ToUpper(newid)

	return newid
}

func mapInt(number int, sourceRange int, destRange int) int {
	return number * destRange / sourceRange
}
