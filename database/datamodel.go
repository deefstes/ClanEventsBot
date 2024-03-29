package database

import (
	"fmt"
	"time"
)

// ClanConfig holds configuration information for the clan's instance of the bot
type ClanConfig struct {
	DefaultChannel string `bson:"defaultChannel" json:"defaultChannel"`
	//InsultProbability float64 `bson:"insultProbability" json:"insultProbability"`
	InsultInterval int64   `bson:"insultInterval" json:"insultInterval"`
	InsultRndFact  float64 `bson:"insultRndFact" json:"insultRndFact"`
}

// Guild holds Discord server information
type Guild struct {
	ID   string `bson:"discordId" json:"discordId"`
	Name string `bson:"name" json:"name"`
}

// TimeZone holds information pertaining to a time zone
type TimeZone struct {
	Abbrev   string `bson:"abbrev" json:"abbrev"`
	Location string `bson:"location" json:"location"`
	Emoji    string `bson:"emoji,omitempty" json:"emoji,omitempty"`
}

type ServerRoleTimeZone struct {
	RoleName string `bson:"roleName" json:"roleName"`
	Abbrev   string `bson:"abbrev" json:"abbrev"`
}

// ClanEvent holds information pertaining to an event
type ClanEvent struct {
	EventID      string     `bson:"eventId" json:"eventId"`
	Creator      ClanUser   `bson:"creator" json:"creator"`
	DateTime     time.Time  `bson:"dateTime" json:"dateTime"`
	TimeZone     string     `bson:"timeZone,omitempty" json:"timeZone,omitempty"`
	Duration     int        `bson:"duration" json:"duration"`
	Name         string     `bson:"name" json:"name"`
	Description  string     `bson:"description" json:"description"`
	TeamSize     int        `bson:"teamSize" json:"teamSize"`
	Full         bool       `bson:"full" json:"full"`
	Participants []ClanUser `bson:"participants" json:"participants"`
	Reserves     []ClanUser `bson:"reserves" json:"reserves"`
	Archived     bool       `bson:"archived" json:"archived"`
}

// ClanUser holds information pertaining to a Discord user
type ClanUser struct {
	UserName          string    `bson:"userName" json:"userName"`
	Nickname          string    `bson:"nickname,omitempty" json:"nickname,omitempty"`
	MentionDeprecated string    `bson:"mention,omitempty" json:"mention,omitempty"`
	UserID            string    `bson:"userId" json:"userId"`
	DateTime          time.Time `bson:"dateTime" json:"dateTime"`
}

// DisplayName returns the user's nickname or username if the former is empty
func (c ClanUser) DisplayName() string {
	if c.Nickname != "" {
		return c.Nickname
	}

	return c.UserName
}

// Mention returns the user's Discord mention name
func (c ClanUser) Mention() string {
	if c.MentionDeprecated != "" {
		return c.MentionDeprecated
	}

	return fmt.Sprintf("<@%s>", c.UserID)
}
