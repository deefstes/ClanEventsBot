package main

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func PingPong(s *discordgo.Session, m *discordgo.MessageCreate, command []string) {
	if command[0] == "ping" {
		s.ChannelMessageSend(m.ChannelID, "pong")
	}
	if command[0] == "pong" {
		s.ChannelMessageSend(m.ChannelID, "ping")
	}
}

func ListEvents(s *discordgo.Session, m *discordgo.MessageCreate, command []string) {
	ct := time.Now()
	reply := fmt.Sprintf("%s", ct)
	s.ChannelMessageSend(m.ChannelID, reply)
}
