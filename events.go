package main

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type ClanEvent struct {
	ObjectId     bson.ObjectId `bson:"_id,omitempty" json:"-"`
	EventId      string        `bson:"eventId" json:"eventId"`
	Creator      ClanUser      `bson:"creator" json:"creator"`
	DateTime     time.Time     `bson:"dateTime" json:"dateTime"`
	Duration     int           `bson:"duration" json:"duration"`
	Name         string        `bson:"name" json:"name"`
	Description  string        `bson:"description" json:"description"`
	TeamSize     int           `bson:"teamSize" json:"teamSize"`
	Full         bool          `bson:"full" json:"full"`
	Participants []ClanUser    `bson:"participants" json:"participants"`
	Reserves     []ClanUser    `bson:"reserves" json:"reserves"`
}

type ClanUser struct {
	UserName string    `bson:"userName" json:"userName"`
	Mention  string    `bson:"mention" json:"mention"`
	DateTime time.Time `bson:"dateTime" json:"dateTime"`
}

type Person struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Name      string
	Phone     string
	Timestamp time.Time
}
