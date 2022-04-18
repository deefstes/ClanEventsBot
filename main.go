package main

import (
	"ClanEventsBot/database"
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
)

var (
	buildNumber     string
	liveTime        time.Time
	config          Configuration
	db              *database.Database
	discordSession  *discordgo.Session
	defaultLocation *time.Location
	guildVars       map[string]*GuildVars
	ErrNoRecords    = errors.New("no records")
)

type GuildVars struct {
	guild          database.Guild
	impersonated   database.ClanUser
	timezones      []database.TimeZone
	tzByAbbr       map[string]database.TimeZone
	tzByEmoji      map[string]database.TimeZone
	escrowEvents   map[string]DevelopingEvent
	defaultChannel string
	insultInterval int64
	insultRndFact  float64
	insultTicker   *time.Ticker
}

func (g *GuildVars) startInsultTimer() {
	g.stopInsultTimer()
	if g.insultInterval == 0 {
		return
	}

	d := time.Duration(g.insultInterval) * time.Minute
	dd := time.Duration(float64(d) * g.insultRndFact)
	fmt.Println("starting insult timer on guild", g.guild.ID, "to fire every", d, "±", dd)
	g.insultTicker = time.NewTicker(time.Duration(g.insultInterval) * time.Minute)
	// defer g.insultTicker.Stop()
	go func() {
		max := int(float64(g.insultInterval) * 60 * (1 + g.insultRndFact))
		min := int(float64(g.insultInterval) * 60 * (1 - g.insultRndFact))
		if min == 0 {
			min = max
		}
		for range g.insultTicker.C {
			// Set new random interval within specified bounds
			dur := time.Duration(max) * time.Second
			if max != min {
				dur = time.Duration(rand.Intn(max-min)+min) * time.Second
			}
			g.insultTicker.Reset(dur)
			fmt.Println("resetting insult timer on guild", g.guild.ID, "to fire after", dur)
			deliverInsult(g)
		}
	}()
}

func (g *GuildVars) stopInsultTimer() {
	if g.insultTicker != nil {
		fmt.Println("stopping insult timer on guild", g.guild.ID)
		g.insultTicker.Stop()
	}
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

	fmt.Printf("ClanEventsBot (build number %s)", buildNumber)
	if config.DebugLevel > 0 {
		fmt.Printf(" with DebugLevel=%d\r\n", config.DebugLevel)
	} else {
		fmt.Println()
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
	db, err = database.NewDatabase(config.MongoDB)
	if err != nil {
		fmt.Println("FATAL", "connecting to database: %v", err)
		os.Exit(1)
	}
	defer db.Close()

	guilds, err := db.GetGuilds()
	if len(guilds) == 0 {
		fmt.Printf("No registered guilds\r\n")
	}

	for _, guild := range guilds {
		guildVars[guild.ID] = GetGuildVars(guild)
	}

	// Create a new Discord session using the provided bot token.
	discordSession, err = discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("ERROR", "creating Discord session,", err)
		return
	}
	defer discordSession.Close()

	discordSession.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuilds | discordgo.IntentsGuildMembers | discordgo.IntentsGuildMessages | discordgo.IntentsGuildMessageReactions)

	// Set up a svcTicker that triggers a service routine every n minutes
	fmt.Println("starting service routine to fire every", config.ServiceTimer, "minutes")
	svcTicker := time.NewTicker(time.Duration(config.ServiceTimer) * time.Minute)
	defer svcTicker.Stop()
	go func() {
		for range svcTicker.C {
			serviceRoutine()
		}
	}()

	// Set up a ticker for each registered guild to deliver insults
	for _, g := range guildVars {
		g.startInsultTimer()
		defer g.stopInsultTimer()
		// if g.insultInterval == 0 {
		// 	continue
		// }

		// d := time.Duration(g.insultInterval) * time.Minute
		// dd := time.Duration(float64(d) * g.insultRndFact)
		// fmt.Println("starting insult timer on guild", g.guild.ID, "to fire every", d, "±", dd)
		// g.insultTicker = time.NewTicker(time.Duration(g.insultInterval) * time.Minute)
		// defer g.insultTicker.Stop()
		// go func(gv *GuildVars) {
		// 	min := int(gv.insultInterval * (1 - gv.insultRndFact))
		// 	max := int(gv.insultInterval * (1 + gv.insultRndFact))
		// 	for range gv.insultTicker.C {
		// 		// Set new random interval within specified bounds
		// 		dur := time.Duration(max) * time.Minute
		// 		if max != min {
		// 			dur = time.Duration(rand.Intn(max-min)+min) * time.Minute
		// 		}
		// 		gv.insultTicker.Reset(dur)
		// 		deliverInsult(gv)
		// 	}
		// }(g)
	}

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
	fmt.Println("Press CTRL-C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func GetGuildVars(g database.Guild) *GuildVars {
	ee := make(map[string]DevelopingEvent)

	tzs, err := db.GetTimeZones(g.ID)
	if err != nil {
		fmt.Println("ERROR", err)
		return nil
	}
	tzBA, tzBE := constructTZMaps(tzs)

	conf, err := db.GetClanConfig(g.ID)
	if err != nil {
		fmt.Println("ERROR", err)
		return nil
	}

	return &GuildVars{
		guild:          g,
		timezones:      tzs,
		tzByAbbr:       tzBA,
		tzByEmoji:      tzBE,
		escrowEvents:   ee,
		defaultChannel: conf.DefaultChannel,
		insultInterval: conf.InsultInterval,
		insultRndFact:  conf.InsultRndFact,
	}
}

func constructTZMaps(tzs []database.TimeZone) (tzBA map[string]database.TimeZone, tzBE map[string]database.TimeZone) {
	tzBA = make(map[string]database.TimeZone)
	tzBE = make(map[string]database.TimeZone)

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
	} else if strings.HasPrefix(command, "remindnaughtylist ") {
		RemindNaughty(guild, s, m, commandElements)
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
	guilds, err := db.GetGuilds()
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}
	if len(guilds) == 0 {
		fmt.Println("No registered guilds")
		return
	}

	for _, guild := range guilds {
		serviceGuild(guild)
	}
}

func serviceGuild(guild database.Guild) {
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
	if err := db.ArchiveEvents(guildID); err != nil {
		fmt.Println("ERROR", err)
	}
}

func deliverInsult(g *GuildVars) {
	insultee, err := getInsultee(g.guild.ID)
	if err == ErrNoRecords {
		return
	} else if err != nil {
		fmt.Println("ERROR", "delivering insult on guild %s: %v", g.guild.ID, err)
		return
	}
	message := getInsult(insultee.Mention())
	if message != "" {
		sendMessage(g.defaultChannel, message)
	}
}

func getInsultee(guildID string) (*database.ClanUser, error) {
	insultees, err := db.GetNaughtyList(guildID)
	if err != nil {
		return nil, fmt.Errorf("database: %w", err)
	}
	if len(insultees) == 0 {
		return nil, ErrNoRecords
	}

	index := rand.Intn(len(insultees))
	return &insultees[index], nil
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
