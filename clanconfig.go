package main

import (
	"gopkg.in/mgo.v2/bson"
)

// ClanConfig holds configuration information for the clan's instance of the bot
type ClanConfig struct {
	ObjectID          bson.ObjectId `bson:"_id,omitempty" json:"-"`
	DefaultChannel    string        `bson:"defaultChannel" json:"defaultChannel"`
	InsultProbability float32       `bson:"insultProbability" json:"insultProbability"`
}

// Guild holds Discord server information
type Guild struct {
	ObjectID bson.ObjectId `bson:"_id,omitempty" json:"-"`
	ID       string        `bson:"discordId" json:"discordId"`
	Name     string        `bson:"name" json:"name"`
}
