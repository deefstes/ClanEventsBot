package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ClanConfig holds configuration information for the clan's instance of the bot
type ClanConfig struct {
	ObjectID          primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	DefaultChannel    string             `bson:"defaultChannel" json:"defaultChannel"`
	InsultProbability float64            `bson:"insultProbability" json:"insultProbability"`
}

// Guild holds Discord server information
type Guild struct {
	ObjectID primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	ID       string             `bson:"discordId" json:"discordId"`
	Name     string             `bson:"name" json:"name"`
}
