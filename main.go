package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/kenshaw/baseconv"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	buildNumber    string
	config         Configuration
	mongoSession   *mgo.Session
	impersonated   ClanUser
	timeZone       *time.Location
	discordSession *discordgo.Session
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "%s - FATAL ERROR: %+v\r\n", time.Now().Format("2006-01-02 15:04:05"), r)
		}
	}()

	if buildNumber == "" {
		buildNumber = "N/A"
	}

	var err error

	// Read config file
	config, err = ReadConfig()
	if err != nil {
		fmt.Println("Error reading config file")
		return
	}

	timeZone, _ = time.LoadLocation("Europe/London")

	// Connect to MongoDB
	mongoSession, err = mgo.Dial(config.MongoDB)
	if err != nil {
		panic(err)
	}
	defer mongoSession.Close()

	// Create a new Discord session using the provided bot token.
	discordSession, err = discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	defer discordSession.Close()

	// Set up a ticker that triggers a service routine every n minutes
	ticker := time.NewTicker(time.Minute * config.ServiceTimer)
	defer ticker.Stop()
	go func() {
		for {
			<-ticker.C
			serviceRoutine()
		}
	}()

	// Register the messageCreate func as a callback for MessageCreate events.
	discordSession.AddHandler(messageCreate)

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

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if message.Author.ID == session.State.User.ID {
		return
	}

	// If messages is not a command, bail out
	if !strings.HasPrefix(message.Content, config.CommandPrefix) {
		return
	}
	command := strings.TrimPrefix(message.Content, config.CommandPrefix)
	commandElements := getArgs(command)

	// If message does not contain guild information, bail out
	guild := getGuild(session, message)
	if guild == nil {
		return
	}

	if config.DebugLevel > 0 {
		fmt.Printf("Guild=%s, Author=%s(%s), Command=%s\r\n", guild.Name, message.Author.Username, message.Author.ID, command)
	}

	if strings.HasPrefix(command, "help") {
		BotHelp(guild, session, message, commandElements)
	}

	if strings.HasPrefix(command, "listevents") {
		ListEvents(guild, session, message, commandElements)
	}

	if strings.HasPrefix(command, "newevent") {
		NewEvent(guild, session, message, commandElements)
	}

	if strings.HasPrefix(command, "cancelevent") {
		CancelEvent(guild, session, message, commandElements)
	}

	if strings.HasPrefix(command, "signup") {
		Signup(guild, session, message, commandElements)
	}

	if strings.HasPrefix(command, "leave") {
		Leave(guild, session, message, commandElements)
	}

	if strings.HasPrefix(command, "impersonate") {
		Impersonate(guild, session, message, commandElements)
	}

	if strings.HasPrefix(command, "unimpersonate") {
		Unimpersonate(guild, session, message, commandElements)
	}

	if strings.HasPrefix(command, "details") {
		Details(guild, session, message, commandElements)
	}

	if strings.HasPrefix(command, "test") {
		Test(guild, session, message, commandElements)
	}

	if strings.HasPrefix(command, "wisdom") {
		Wisdom(guild, session, message, commandElements)
	}

	if strings.HasPrefix(command, "addnaughtylist") {
		AddNaughty(guild, session, message, commandElements)
	}

	if strings.HasPrefix(command, "removenaughtylist") {
		RemoveNaughty(guild, session, message, commandElements)
	}

	if strings.HasPrefix(command, "addserver") {
		AddServer(guild, session, message, commandElements)
	}
}

func sendMessage(channelID string, message string) {
	discordSession.ChannelMessageSend(channelID, message)
}

func serviceRoutine() {
	c := mongoSession.DB("ClanEvents").C("Guilds")

	var guilds []Guild

	err := c.Find(bson.M{}).All(&guilds)
	if err != nil {
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

	c := mongoSession.DB(fmt.Sprintf("ClanEvents%s", guildID)).C("Events")

	var results []ClanEvent
	err := c.Find(filter).Sort("dateTime").All(&results)
	if err != nil {
		fmt.Printf("Error archiving events on guild %s\r\n", guildID)
		return
	}

	for _, event := range results {
		upsertfilter := bson.M{"eventId": event.EventID}
		event.Archived = true
		event.EventID = fmt.Sprintf("%s_%s", time.Now().Format("060102150405"), event.EventID)
		_, err := c.Upsert(upsertfilter, event)
		if err != nil {
			fmt.Printf("Error archiving event %s on guild %s\r\n", event.EventID, guildID)
			return
		}
	}
}

func deliverInsult(guildID string) {
	// Find default channel and insultee in DB
	c := mongoSession.DB(fmt.Sprintf("ClanEvents%s", guildID)).C("Config")
	var config ClanConfig
	err := c.Find(bson.M{}).One(&config)
	if err != nil {
		fmt.Printf("Error delivering insult on guild %s\r\n", guildID)
		return
	}

	prob := rand.Float32()
	if prob <= config.InsultProbability {
		insultee, err := getInsultee(guildID)
		if err != nil {
			fmt.Printf("Error delivering insult on guild %s\r\n", guildID)
			return
		}
		message := getInsult(insultee.Mention())
		sendMessage(config.DefaultChannel, message)
	}
}

func getInsultee(guildID string) (ClanUser, error) {
	c := mongoSession.DB(fmt.Sprintf("ClanEvents%s", guildID)).C("NaughtyList")

	var insultees []ClanUser
	var insultee ClanUser

	err := c.Find(bson.M{}).All(&insultees)
	if err != nil {
		fmt.Printf("Error finding someone on guild %s to insult\r\n", guildID)
		return insultee, errors.New("No insultee found")
	}

	if len(insultees) == 0 {
		fmt.Printf("Error finding someone on guild %s to insult\r\n", guildID)
		return insultee, errors.New("No insultee found")
	}

	index := rand.Intn(len(insultees))
	insultee = insultees[index]

	return insultee, nil
}

func getInsult(mention string) string {
	canon := []string{
		fmt.Sprintf("%s is a stuck up, half-witted, scruffy-looking… Nerf herder!", mention),
		fmt.Sprintf("%s's parents are living proof that two wrongs don't make a right.", mention),
		fmt.Sprintf("EventsBot farts in %s's general direction.", mention),
		fmt.Sprintf("%s's brain is so minute that if a hungry cannibal cracked the head open, there wouldn't be enough to cover a small water biscuit.", mention),
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
	}

	index := rand.Intn(len(canon))
	return canon[index]
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
