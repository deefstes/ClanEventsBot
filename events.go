package main

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// ClanEvent holds information pertaining to an event
type ClanEvent struct {
	ObjectID     bson.ObjectId `bson:"_id,omitempty" json:"-"`
	EventID      string        `bson:"eventId" json:"eventId"`
	Creator      ClanUser      `bson:"creator" json:"creator"`
	DateTime     time.Time     `bson:"dateTime" json:"dateTime"`
	Duration     int           `bson:"duration" json:"duration"`
	Name         string        `bson:"name" json:"name"`
	Description  string        `bson:"description" json:"description"`
	TeamSize     int           `bson:"teamSize" json:"teamSize"`
	Full         bool          `bson:"full" json:"full"`
	Participants []ClanUser    `bson:"participants" json:"participants"`
	Reserves     []ClanUser    `bson:"reserves" json:"reserves"`
	Archived     bool          `bson:"archived" json:"archived"`
}

// ClanUser holds information pertaining to a Discord user
type ClanUser struct {
	UserName string    `bson:"userName" json:"userName"`
	Mention  string    `bson:"mention" json:"mention"`
	DateTime time.Time `bson:"dateTime" json:"dateTime"`
}
