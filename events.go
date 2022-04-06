package main

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ClanEvent holds information pertaining to an event
type ClanEvent struct {
	ObjectID     primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	EventID      string             `bson:"eventId" json:"eventId"`
	Creator      ClanUser           `bson:"creator" json:"creator"`
	DateTime     time.Time          `bson:"dateTime" json:"dateTime"`
	TimeZone     string             `bson:"timeZone,omitempty" json:"timeZone,omitempty"`
	Duration     int                `bson:"duration" json:"duration"`
	Name         string             `bson:"name" json:"name"`
	Description  string             `bson:"description" json:"description"`
	TeamSize     int                `bson:"teamSize" json:"teamSize"`
	Full         bool               `bson:"full" json:"full"`
	Participants []ClanUser         `bson:"participants" json:"participants"`
	Reserves     []ClanUser         `bson:"reserves" json:"reserves"`
	Archived     bool               `bson:"archived" json:"archived"`
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
