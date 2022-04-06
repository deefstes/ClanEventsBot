package main

import "go.mongodb.org/mongo-driver/bson/primitive"

// TimeZone holds information pertaining to a time zone
type TimeZone struct {
	ObjectID primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	Abbrev   string             `bson:"abbrev" json:"abbrev"`
	Location string             `bson:"location" json:"location"`
	Emoji    string             `bson:"emoji,omitempty" json:"emoji,omitempty"`
}

type ServerRoleTimeZone struct {
	ObjectID primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	RoleName string             `bson:"roleName" json:"roleName"`
	Abbrev   string             `bson:"abbrev" json:"abbrev"`
}
