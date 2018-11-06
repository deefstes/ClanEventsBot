package main

import (
	"gopkg.in/mgo.v2/bson"
)

// TimeZone holds information pertaining to a time zone
type TimeZone struct {
	ObjectID bson.ObjectId `bson:"_id,omitempty" json:"-"`
	Abbrev   string        `bson:"abbrev" json:"abbrev"`
	Location string        `bson:"location" json:"location"`
}

type ServerRoleTimeZone struct {
	ObjectID bson.ObjectId `bson:"_id,omitempty" json:"-"`
	RoleName string        `bson:"roleName" json:"roleName"`
	Abbrev   string        `bson:"abbrev" json:"abbrev"`
}
