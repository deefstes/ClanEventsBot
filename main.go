package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	config Configuration
)

func main() {
	var ok bool
	config, ok = ReadConfig()
	if !ok {
		fmt.Println("Error reading config file")
		return
	}

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
	commandElements := strings.Fields(command)

	if strings.HasPrefix(command, "list") {
		ListEvents(s, m, commandElements)
	}

	if strings.HasPrefix(command, "ping") {
		PingPong(s, m, commandElements)
	}

	if strings.HasPrefix(command, "pong") {
		PingPong(s, m, commandElements)
	}
}
