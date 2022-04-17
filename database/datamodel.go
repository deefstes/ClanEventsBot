package database

import (
	"fmt"
	"time"
)

// ClanConfig holds configuration information for the clan's instance of the bot
type ClanConfig struct {
	//ObjectID          primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	DefaultChannel    string  `bson:"defaultChannel" json:"defaultChannel"`
	InsultProbability float64 `bson:"insultProbability" json:"insultProbability"`
}

// Guild holds Discord server information
type Guild struct {
	//ObjectID primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	ID   string `bson:"discordId" json:"discordId"`
	Name string `bson:"name" json:"name"`
}

// TimeZone holds information pertaining to a time zone
type TimeZone struct {
	//ObjectID primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	Abbrev   string `bson:"abbrev" json:"abbrev"`
	Location string `bson:"location" json:"location"`
	Emoji    string `bson:"emoji,omitempty" json:"emoji,omitempty"`
}

type ServerRoleTimeZone struct {
	//ObjectID primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	RoleName string `bson:"roleName" json:"roleName"`
	Abbrev   string `bson:"abbrev" json:"abbrev"`
}

// ClanEvent holds information pertaining to an event
type ClanEvent struct {
	//ObjectID     primitive.ObjectID `bson:"_id,omitempty" json:"-"`
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
	UserName           string    `bson:"userName" json:"userName"`
	Nickname           string    `bson:"nickname,omitempty" json:"nickname,omitempty"`
	Mention_deprecated string    `bson:"mention,omitempty" json:"mention,omitempty"`
	UserID             string    `bson:"userId" json:"userId"`
	DateTime           time.Time `bson:"dateTime" json:"dateTime"`
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
	if c.Mention_deprecated != "" {
		return c.Mention_deprecated
	}

	return fmt.Sprintf("<@%s>", c.UserID)
}
