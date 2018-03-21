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
	config, ok := ReadConfig()
	if !ok {
		fmt.Println("Error reading config file")
		return
	}
	fmt.Printf("Token: %s\r\n", config.Token)
	fmt.Printf("Command Prefix: %s\r\n", config.CommandPrefix)

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
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	fmt.Printf("Token: %s\r\n", config.Token)
	fmt.Printf("Prefix: %s\r\n", config.CommandPrefix)
	if strings.HasPrefix(m.Content, config.CommandPrefix+"list") {
		ListEvents(s, m)
	}

	// If the message is "ping" reply with "pong"
	if m.Content == config.CommandPrefix+"ping" {
		s.ChannelMessageSend(m.ChannelID, "pong")
	}

	// If the message is "pong" reply with "pong"
	if m.Content == config.CommandPrefix+"pong" {
		s.ChannelMessageSend(m.ChannelID, "ping")
	}
}
