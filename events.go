package main

import "time"

type ClanEvent struct {
	id       int
	creator  ClanUser
	dateTime time.Time
}

type ClanUser struct {
	userName string
	mention  string
	dateTime time.Time
}
