package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func ListEvents(s *discordgo.Session, m *discordgo.MessageCreate) {
	words := strings.Fields(m.Content)
	for _, word := range words {
		fmt.Println(word)
	}
}
