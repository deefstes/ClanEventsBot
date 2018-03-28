package main

import (
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"gopkg.in/mgo.v2"
)

var (
	config       Configuration
	mongoSession *mgo.Session
	impersonated ClanUser
	timeZone     *time.Location
)

func main() {
	var err error
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
	dg, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("ClanEventsBot is now running")
	fmt.Println("Press CTRL-C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// If messages is not a command, bail out
	if !strings.HasPrefix(m.Content, config.CommandPrefix) {
		return
	}
	command := strings.TrimPrefix(m.Content, config.CommandPrefix)
	commandElements := getArgs(command)

	if strings.HasPrefix(command, "help") {
		BotHelp(s, m, commandElements)
	}

	if strings.HasPrefix(command, "listevents") {
		ListEvents(s, m, commandElements)
	}

	if strings.HasPrefix(command, "newevent") {
		NewEvent(s, m, commandElements)
	}

	if strings.HasPrefix(command, "cancelevent") {
		CancelEvent(s, m, commandElements)
	}

	if strings.HasPrefix(command, "signup") {
		Signup(s, m, commandElements)
	}

	if strings.HasPrefix(command, "leave") {
		Leave(s, m, commandElements)
	}

	if strings.HasPrefix(command, "impersonate") {
		Impersonate(s, m, commandElements)
	}

	if strings.HasPrefix(command, "unimpersonate") {
		Unimpersonate(s, m, commandElements)
	}

	if strings.HasPrefix(command, "details") {
		Details(s, m, commandElements)
	}

	if strings.HasPrefix(command, "test") {
		Test(s, m, commandElements)
	}
}

func getArgs(s string) []string {
	re := regexp.MustCompile("\".+?\"|'.+?'|\\S+")
	args := re.FindAllString(s, -1)
	for i := 0; i < len(args); i++ {
		args[i] = strings.TrimPrefix(args[i], "\"")
		args[i] = strings.TrimSuffix(args[i], "\"")
		args[i] = strings.TrimPrefix(args[i], "'")
		args[i] = strings.TrimSuffix(args[i], "'")
	}
	return args
}
