package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/kenshaw/baseconv"
	"gopkg.in/mgo.v2/bson"
)

func PingPong(s *discordgo.Session, m *discordgo.MessageCreate, command []string) {
	if command[0] == "ping" {
		s.ChannelMessageSend(m.ChannelID, "pong")
	}
	if command[0] == "pong" {
		s.ChannelMessageSend(m.ChannelID, "ping")
	}
}

/// ~listevents
/// ~listevents @username
/// ~listevents date
/// ~listevents date @username
func ListEvents(s *discordgo.Session, m *discordgo.MessageCreate, command []string) {
	message := ""

	// Test for correct number of arguments
	if len(command) > 3 {
		message = fmt.Sprintf("Whoah, not so sure about those arguments. EventsBot is confused :thinking:")
		message = fmt.Sprintf("%s\r\nFor help with listing events, type the following:\r\n```%shelp listevents```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	var specdate time.Time
	listuser := "all"

	// Check first argument
	if len(command) > 1 {
		if IsUser(command[1]) {
			listuser = m.Mentions[0].Username
		} else if IsDate(command[1]) {
			specdate, _ = time.ParseInLocation("2006-01-02", command[1], timeZone)
		} else {
			message = fmt.Sprintf("Whoah, not so sure about those arguments. EventsBot is confused :confounded:")
			message = fmt.Sprintf("%s\r\nFor help with listing events, type the following:\r\n```%shelp listevents```", message, config.CommandPrefix)
			s.ChannelMessageSend(m.ChannelID, message)
			return
		}
	}

	// Check second argument
	if len(command) > 2 {
		if IsUser(command[2]) {
			listuser = m.Mentions[0].Username
		} else {
			message = fmt.Sprintf("Whoah, not so sure about those arguments. EventsBot is confused :anguished:")
			message = fmt.Sprintf("%s\r\nFor help with listing events, type the following:\r\n```%shelp listevents```", message, config.CommandPrefix)
			s.ChannelMessageSend(m.ChannelID, message)
			return
		}
	}

	filter := bson.M{}
	filter["dateTime"] = bson.M{
		"$gte": time.Now(),
	}
	if listuser != "all" {
		filter["participants.userName"] = listuser
	}
	if !specdate.IsZero() {
		filter["dateTime"] = bson.M{
			"$gte": specdate,
			"$lt":  specdate.AddDate(0, 0, 1),
		}
	}

	c := mongoSession.DB("ClanEvents").C("Events")

	var results []ClanEvent
	err := c.Find(filter).Sort("dateTime").All(&results)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, ":scream::scream::scream:Something very weird happened when trying to read the events. Sorry but EventsBot has no answers for you :cry:")
		return
	}

	var reply string
	if specdate.IsZero() {
		reply = "Upcoming clan events"
	} else {
		reply = fmt.Sprintf("Clan events on %s", specdate.Format("2006-01-02"))
	}
	reply = fmt.Sprintf("%s for %s\r\n", reply, listuser)
	if len(results) == 0 {
		reply = fmt.Sprintf("%sZip. Nothing. Nada.\r\nWhat nonsense is this? EventsBot does not approve :frowning2:", reply)
	} else {
		reply = fmt.Sprintf("%s```", reply)
		for _, event := range results {
			reply = fmt.Sprintf("%s[%s] %s - %s -", reply, event.EventId, event.DateTime.In(timeZone).Format("2006-01-02 15:04"), event.Name)
			for _, participant := range event.Participants {
				reply = fmt.Sprintf("%s %s,", reply, participant.UserName)
			}
			reply = fmt.Sprintf("%s\r\n\r\n", strings.TrimSuffix(reply, ","))
		}
		reply = fmt.Sprintf("%s```", reply)
	}

	s.ChannelMessageSend(m.ChannelID, reply)
}

/// ~details eventId
func Details(s *discordgo.Session, m *discordgo.MessageCreate, command []string) {
	message := ""

	// Test for correct number of arguments
	if len(command) != 2 {
		message = fmt.Sprintf("Whoah, not so sure about those arguments. EventsBot is confused :thinking:")
		message = fmt.Sprintf("%s\r\nFor help with getting the details of an event, type the following:\r\n```%shelp details```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	// Find event in DB
	c := mongoSession.DB("ClanEvents").C("Events")

	var event ClanEvent
	err := c.Find(bson.M{"eventId": command[1]}).One(&event)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("EventsBot could find no such event. Are you sure you got that Event ID of %s right? Them's finicky numbers. Remember, they're case sensitive :grimacing:", command[1]))
		return
	}

	message = fmt.Sprintf("**EventID:** %s", event.EventId)
	message = fmt.Sprintf("%s\r\n**Creator:** %s", message, event.Creator.Mention)
	message = fmt.Sprintf("%s\r\n**Date:** %s", message, event.DateTime.In(timeZone).Format("2006-01-02"))
	message = fmt.Sprintf("%s\r\n**Time:** %s for %d hours", message, event.DateTime.In(timeZone).Format("15:04"), event.Duration)
	message = fmt.Sprintf("%s\r\n**Name:** %s", message, event.Name)
	message = fmt.Sprintf("%s\r\n**Description:** %s", message, event.Description)
	message = fmt.Sprintf("%s\r\n**Team Size:** %d of %d", message, len(event.Participants), event.TeamSize)
	if len(event.Participants) > 0 {
		message = fmt.Sprintf("%s\r\n**Participants:**", message)
		for _, participant := range event.Participants {
			message = fmt.Sprintf("%s\r\n   -  %s", message, participant.Mention)
		}
	}
	if len(event.Reserves) > 0 {
		message = fmt.Sprintf("%s\r\n**Reserves:**", message)
		for _, reserve := range event.Reserves {
			message = fmt.Sprintf("%s\r\n    - %s", message, reserve.Mention)
		}
	}

	s.ChannelMessageSend(m.ChannelID, message)
}

/// ~newevent Date Time Duration Name Description Size
func NewEvent(s *discordgo.Session, m *discordgo.MessageCreate, command []string) {
	message := ""

	// Test for correct number of arguments
	if len(command) != 7 {
		message = fmt.Sprintf("Whoah, not so sure about those arguments. EventsBot is confused :thinking:")
		message = fmt.Sprintf("%s\r\nFor help with creating a new event, type the following:\r\n```%shelp newevent```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	// Test for date and time arguments
	datetime := fmt.Sprintf("%s %s", command[1], command[2])
	dt, err := time.ParseInLocation("2006-01-02 15:04", datetime, timeZone)
	if err != nil {
		message = fmt.Sprintf("Whoah, not so sure about that date and time (%s). EventsBot is confused :thinking:", datetime)
		message = fmt.Sprintf("%s\r\nFor help with creating a new event, type the following:\r\n```%shelp newevent```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}
	if dt.Before(time.Now()) {
		message = fmt.Sprintf("Are you trying to create an event in the past? EventsBot has lost his flux capacitor :robot:")
		message = fmt.Sprintf("%s\r\nFor help with creating a new event, type the following:\r\n```%shelp newevent```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	// Test for duration
	duration, err := strconv.Atoi(command[3])
	if err != nil {
		message = fmt.Sprintf("What kind of a duration is %s? EventsBot needs a vacation of %s weeks :beach:", command[3], command[3])
		message = fmt.Sprintf("%s\r\nFor help with creating a new event, type the following:\r\n```%shelp newevent```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	// Test for size
	teamSize, err := strconv.Atoi(command[6])
	if err != nil {
		message = fmt.Sprintf("How many players you say? %s? EventsBot wouldn't do that if he were you :speak_no_evil:", command[6])
		message = fmt.Sprintf("%s\r\nFor help with creating a new event, type the following:\r\n```%shelp newevent```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	newid, _ := baseconv.Encode62FromDec(time.Now().Format("050415020106")) // Convert the current DateTime in reverse order of significance (ssmmHHDDMMYY) to Base62
	curUser := impersonated
	if curUser.UserName == "" {
		curUser = ClanUser{
			UserName: m.Author.Username,
			Mention:  m.Author.Mention(),
			DateTime: time.Now(),
		}
	}
	newEvent := ClanEvent{
		EventId:     newid,
		Creator:     curUser,
		DateTime:    dt,
		Duration:    duration,
		Name:        command[4],
		Description: command[5],
		TeamSize:    teamSize,
		Full:        false,
	}

	c := mongoSession.DB("ClanEvents").C("Events")
	err = c.Insert(newEvent)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, ":scream::scream::scream:Something very weird happened when trying to create this event. Sorry but EventsBot has no answers for you :cry:")
		return
	}

	message = fmt.Sprintf("Woohoo! A new event has been created by %s. EventsBot is most pleased :ok_hand:", newEvent.Creator.Mention)
	message = fmt.Sprintf("%s\r\nEvent ID: **%s**", message, newEvent.EventId)
	message = fmt.Sprintf("%s\r\n\r\nTo sign up for this event, type the following:", message)
	message = fmt.Sprintf("%s\r\n```%ssignup %s```", message, config.CommandPrefix, newEvent.EventId)
	s.ChannelMessageSend(m.ChannelID, message)
}

/// ~cancelevent eventId
func CancelEvent(s *discordgo.Session, m *discordgo.MessageCreate, command []string) {
	message := ""

	// Test for correct number of arguments
	if len(command) != 2 {
		message = fmt.Sprintf("Whoah, not so sure about those arguments. EventsBot is confused :thinking:")
		message = fmt.Sprintf("%s\r\nFor help with cancelling an event, type the following:\r\n```%shelp cancelevent```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	curUser := impersonated
	curUser.DateTime = time.Now()
	if curUser.UserName == "" {
		curUser = ClanUser{
			UserName: m.Author.Username,
			Mention:  m.Author.Mention(),
			DateTime: time.Now(),
		}
	}

	// Find event in DB
	c := mongoSession.DB("ClanEvents").C("Events")

	var event ClanEvent
	err := c.Find(bson.M{"eventId": command[1]}).One(&event)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("EventsBot could find no such event. Are you sure you got that Event ID of %s right? Them's finicky numbers. Remember, they're case sensitive :grimacing:", command[1]))
		return
	}

	// Check that user has permissions
	allowed := false
	if event.Creator.UserName == curUser.UserName {
		allowed = true
	} else if HasRole(s, m, "EventsBotAdmin") {
		allowed = true
	}

	if !allowed {
		message = fmt.Sprintf("Yo yo yo. Back up a second dude. You don't have permissions to cancel this event.\r\nEventsBot will not stand for this :point_up:")
		message = fmt.Sprintf("%s\r\nFor help with cancelling events, type the following:\r\n```%shelp cancelevent```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	// Delete record
	err = c.Remove(bson.M{"eventId": command[1]})
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, ":scream::scream::scream:Something very weird happened when trying to create this event. Sorry but EventsBot has no answers for you :cry:")
		return
	}

	message = fmt.Sprintf("Tragedy! %s's event, %s, has been cancelled.", event.Creator.Mention, event.Name)
	message = fmt.Sprintf("%s\r\n\r\n\"We don't have commit. Repeat. We are decommissioning the committal of the launch. It is now a negatory launch phase. We are in a no fly, no go phase. That is a November Gorgon phase, of non-flying. And we're gonna say 'goodnight, thank you, good work, over and out'\".", message)
	message = fmt.Sprintf("%s\r\n\r\nEventsBot will cry himself to sleep tonight :sob:", message)
	s.ChannelMessageSend(m.ChannelID, message)
}

/// ~signup eventId
/// ~signup eventId @Username
func Signup(s *discordgo.Session, m *discordgo.MessageCreate, command []string) {
	message := ""

	// Test for correct number of arguments
	if len(command) > 3 {
		message = fmt.Sprintf("Whoah, not so sure about those arguments. EventsBot is confused :thinking:")
		message = fmt.Sprintf("%s\r\nFor help with signing up to an event, type the following:\r\n```%shelp signup```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	curUser := impersonated
	curUser.DateTime = time.Now()
	if curUser.UserName == "" {
		curUser = ClanUser{
			UserName: m.Author.Username,
			Mention:  m.Author.Mention(),
			DateTime: time.Now(),
		}
	}

	signupUser := curUser

	// Check first argument
	if len(command) > 2 {
		if IsUser(command[2]) {
			signupUser = ClanUser{
				UserName: m.Mentions[0].Username,
				Mention:  m.Mentions[0].Mention(),
			}
		} else {
			message = fmt.Sprintf("Whoah, not so sure about those arguments. EventsBot is confused :confounded:")
			message = fmt.Sprintf("%s\r\nFor help with signing up to events, type the following:\r\n```%shelp signup```", message, config.CommandPrefix)
			s.ChannelMessageSend(m.ChannelID, message)
			return
		}
	}

	// Find event in DB
	c := mongoSession.DB("ClanEvents").C("Events")

	var event ClanEvent
	err := c.Find(bson.M{"eventId": command[1]}).One(&event)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("EventsBot could find no such event. Are you sure you got that Event ID of %s right? Them's finicky numbers. Remember, they're case sensitive :grimacing:", command[1]))
		return
	}

	// If different user is specified, check that current user has permissions
	if signupUser != curUser {
		allowed := false
		if event.Creator.UserName == curUser.UserName {
			allowed = true
		} else if HasRole(s, m, "EventsBotAdmin") {
			allowed = true
		}

		if !allowed {
			message = fmt.Sprintf("Yo yo yo. Back up a second dude. You don't have permissions to sign other users up to events.\r\nEventsBot will not stand for this :point_up:")
			message = fmt.Sprintf("%s\r\nFor help with signing up to events, type the following:\r\n```%shelp signup```", message, config.CommandPrefix)
			s.ChannelMessageSend(m.ChannelID, message)
			return
		}
	}

	// Check if user is already signed up for this event
	for _, participant := range event.Participants {
		if participant.UserName == signupUser.UserName {
			if signupUser.UserName == curUser.UserName {
				s.ChannelMessageSend(m.ChannelID, "You are already signed up to this event.\r\nEventsBot hasn't got time for your shenanigans :rolling_eyes:")
			} else {
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s is already signed up to this event.\r\nEventsBot hasn't got time for your shenanigans :rolling_eyes:", signupUser.UserName))
			}
			return
		}
	}

	// Check if event is full
	if len(event.Participants) >= event.TeamSize {
		s.ChannelMessageSend(m.ChannelID, "Oh noes! This event is already full :cry:\r\nBut don't worry, EventsBot will put you on the reserves list and notify you if someone leaves.")
		for _, reserve := range event.Reserves {
			if reserve.UserName == curUser.UserName {
				return
			}
		}
		event.Reserves = append(event.Reserves, signupUser)
		err = c.Update(bson.M{"eventId": command[1]}, event)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, ":scream::scream::scream:Something very weird happened when trying to update the event. Sorry but EventsBot has no answers for you :cry:")
			return
		}
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s is now signed up as a reserve for %s's event, %s.\r\nEventsBot approves :thumbsup:", signupUser.Mention, event.Creator.Mention, event.Name))
	} else {
		event.Participants = append(event.Participants, signupUser)
		event.Full = len(event.Participants) >= event.TeamSize
		err = c.Update(bson.M{"eventId": command[1]}, event)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, ":scream::scream::scream:Something very weird happened when trying to update the event. Sorry but EventsBot has no answers for you :cry:")
			return
		}
		message := fmt.Sprintf("%s is now signed up for %s's event, %s.\r\n", signupUser.Mention, event.Creator.Mention, event.Name)
		if event.Full {
			message = fmt.Sprintf("%sThis event is now full. It's all systems go!\r\n", message)
			message = fmt.Sprintf("%sEventsBot definitely approves :thumbsup::thumbsup:", message)
		} else {
			if event.TeamSize-len(event.Participants) == 1 {
				message = fmt.Sprintf("%sThere is one space left\r\n", message)
			} else {
				message = fmt.Sprintf("%sThere are %d spaces left\r\n", message, event.TeamSize-len(event.Participants))
			}
			message = fmt.Sprintf("%sEventsBot approves :thumbsup:", message)
		}
		s.ChannelMessageSend(m.ChannelID, message)
	}
}

/// ~leave eventId
/// ~leave eventId @Username
func Leave(s *discordgo.Session, m *discordgo.MessageCreate, command []string) {
	message := ""

	// Test for correct number of arguments
	if len(command) > 3 || len(command) < 2 {
		message = fmt.Sprintf("Whoah, not so sure about those arguments. EventsBot is confused :thinking:")
		message = fmt.Sprintf("%s\r\nFor help with leaving an event, type the following:\r\n```%shelp leave```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	curUser := impersonated
	curUser.DateTime = time.Now()
	if curUser.UserName == "" {
		curUser = ClanUser{
			UserName: m.Author.Username,
			Mention:  m.Author.Mention(),
			DateTime: time.Now(),
		}
	}

	removeUser := curUser

	// Check first argument
	if len(command) > 2 {
		if IsUser(command[2]) {
			removeUser = ClanUser{
				UserName: m.Mentions[0].Username,
				Mention:  m.Mentions[0].Mention(),
			}
		} else {
			message = fmt.Sprintf("Whoah, not so sure about those arguments. EventsBot is confused :confounded:")
			message = fmt.Sprintf("%s\r\nFor help with signing up to events, type the following:\r\n```%shelp signup```", message, config.CommandPrefix)
			s.ChannelMessageSend(m.ChannelID, message)
			return
		}
	}

	// Find event in DB
	c := mongoSession.DB("ClanEvents").C("Events")

	var event ClanEvent
	err := c.Find(bson.M{"eventId": command[1]}).One(&event)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("EventsBot could find no such event. Are you sure you got that Event ID of %s right? Them's finicky numbers. Remember, they're case sensitive :grimacing:", command[1]))
		return
	}

	// If different user is specified, check that current user has permissions
	if removeUser != curUser {
		allowed := false
		if event.Creator.UserName == curUser.UserName {
			allowed = true
		} else if HasRole(s, m, "EventsBotAdmin") {
			allowed = true
		}

		if !allowed {
			message = fmt.Sprintf("Yo yo yo. Back up a second dude. You don't have permissions to remove other users from events.\r\nEventsBot will not stand for this :point_up:")
			message = fmt.Sprintf("%s\r\nFor help with leaving events, type the following:\r\n```%shelp signup```", message, config.CommandPrefix)
			s.ChannelMessageSend(m.ChannelID, message)
			return
		}
	}

	// Check if user is in fact signed up for this event
	index := -1
	for i, participant := range event.Participants {
		if participant.UserName == removeUser.UserName {
			index = i
		}
		//fmt.Printf("%d02 %d02 %s =?= %s", i, index, )
	}
	if index == -1 {
		if curUser.UserName == removeUser.UserName {
			s.ChannelMessageSend(m.ChannelID, "You are not signed up to this event.\r\nEventsBot does not find your jokes particularly funny :rolling_eyes:")
		} else {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s is not signed up to this event.\r\nEventsBot does not find your jokes particularly funny :rolling_eyes:", removeUser.UserName))
		}
		return
	}

	// Remove participant from event
	event.Participants = append(event.Participants[:index], event.Participants[index+1:]...)
	message = fmt.Sprintf("Well okay then, %s has been removed from %s's event, %s\r\nEventsBot is sad to see you go :disappointed_relieved:", removeUser.Mention, event.Creator.Mention, event.Name)

	// Move first reserve into participants
	if len(event.Reserves) > 0 {
		message = fmt.Sprintf("%s\r\nBut hey! %s is on reserve so we're golden.\r\nEventsBot is relieved :relieved:", message, event.Reserves[0].Mention)
		reserve := ClanUser{
			UserName: event.Reserves[0].UserName,
			Mention:  event.Reserves[0].Mention,
			DateTime: event.Reserves[0].DateTime,
		}
		event.Participants = append(event.Participants, reserve)
		event.Reserves = append(event.Reserves[:0], event.Reserves[0+1:]...)

		err = c.Update(bson.M{"eventId": command[1]}, event)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, ":scream::scream::scream:Something very weird happened when trying to update the event. Sorry but EventsBot has no answers for you :cry:")
			return
		}
	} else {
		//fmt.Printf("%d-%sd = %d\r\n", event.TeamSize, len(event.Participants), event.TeamSize-len(event.Participants))
		if event.TeamSize-len(event.Participants) == 1 {
			message = fmt.Sprintf("%s\r\nThere is now one space left\r\n", message)
		} else {
			message = fmt.Sprintf("%s\r\nThere are now %d spaces left\r\n", message, event.TeamSize-len(event.Participants))
		}
	}

	err = c.Update(bson.M{"eventId": command[1]}, event)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, ":scream::scream::scream:Something very weird happened when trying to update the event. Sorry but EventsBot has no answers for you :cry:")
		return
	}
	s.ChannelMessageSend(m.ChannelID, message)
}

func BotHelp(s *discordgo.Session, m *discordgo.MessageCreate, command []string) {
	if len(command) == 1 {
		command = append(command, "nothing")
	}

	message := fmt.Sprintf("Need some help with the __%s__ command? EventsBot is happy to oblige :nerd:", command[1])

	switch command[1] {
	case "nothing":
		message = "List of EventsBot commands:```"
		message = fmt.Sprintf("%s\r\n    %slistevents", message, config.CommandPrefix)
		message = fmt.Sprintf("%s\r\n    %sdetails", message, config.CommandPrefix)
		message = fmt.Sprintf("%s\r\n    %snewevent", message, config.CommandPrefix)
		message = fmt.Sprintf("%s\r\n    %scancelevent", message, config.CommandPrefix)
		message = fmt.Sprintf("%s\r\n    %ssignup", message, config.CommandPrefix)
		message = fmt.Sprintf("%s\r\n    %sleave", message, config.CommandPrefix)
		message = fmt.Sprintf("%s\r\n    %simpersonate", message, config.CommandPrefix)
		message = fmt.Sprintf("%s\r\n    %sunimpersonate", message, config.CommandPrefix)
		message = fmt.Sprintf("%s```", message)
		message = fmt.Sprintf("%sYou can get help on any of these commands by typing ~help followed by the name of the command", message)
	case "listevents":
		message = fmt.Sprintf("%s\r\nHere's how to get a list of upcoming events:", message)
		message = fmt.Sprintf("%s\r\n```%slistevents Date @Username\r\n", message, config.CommandPrefix)
		message = fmt.Sprintf("%s\r\n       Date: The date for which you want to see events", message)
		message = fmt.Sprintf("%s\r\n  @Username: The Discord user for which you want to see events", message)
		message = fmt.Sprintf("%s\r\n\r\nNote: Both the date and username values are optional. You can specify either, neither or both but then they must be in the order shown above. If you omit the date, you will be shown all upcoming events and if you omit the user you will be shown events for all users.", message)
		message = fmt.Sprintf("%s```", message)
	case "details":
		message = fmt.Sprintf("%s\r\nHere's how to get details for an event:", message)
		message = fmt.Sprintf("%s\r\n```%sdetails EventID\r\n", message, config.CommandPrefix)
		message = fmt.Sprintf("%s\r\n    EventID: That weird looking 7 character identifier that uniquely identifies the event. These values are case sensitive so do take care to get it right. It's your key to participation, enjoyment and a deeper level of zen.", message)
		message = fmt.Sprintf("%s```", message)
	case "newevent":
		message = fmt.Sprintf("%s\r\nHere's how to create a new event:", message)
		message = fmt.Sprintf("%s\r\n```%snewevent Date Time Name Description Size\r\n", message, config.CommandPrefix)
		message = fmt.Sprintf("%s\r\n       Date: In the format YYYY-MM-DD", message)
		message = fmt.Sprintf("%s\r\n       Time: In the format HH:MM (24 hour clock)", message)
		message = fmt.Sprintf("%s\r\n   Duration: Number of hours the event will last", message)
		message = fmt.Sprintf("%s\r\n       Name: A name for your event. Surround it in quotes if it's more than one word", message)
		message = fmt.Sprintf("%s\r\nDescription: A longer description of your event. You totally want to surround this one in quotes", message)
		message = fmt.Sprintf("%s\r\n   TeamSize: Just a number denoting how many players can sign up```", message)
		message = fmt.Sprintf("%s\r\n\r\nHere's an example for you:", message)
		message = fmt.Sprintf("%s\r\n```%snewevent %s 20:00 2 \"Normal Leviathan\" \"Fresh start of Leviathan raid\" 6```", message, config.CommandPrefix, time.Now().Format("2006-01-02"))
		message = fmt.Sprintf("%s\r\nThis will create a 2 hour event to start at 8pm tonight and which will allow 6 people to sign up", message)
	case "cancelevent":
		message = fmt.Sprintf("%s\r\nHere's how to cancel an event:", message)
		message = fmt.Sprintf("%s\r\n```%scancelevent EventID\r\n", message, config.CommandPrefix)
		message = fmt.Sprintf("%s\r\n    EventID: That weird looking 7 character identifier that uniquely identifies the event. These values are case sensitive so do take care to get it right. It's your key to participation, enjoyment and a deeper level of zen.", message)
		message = fmt.Sprintf("%s\r\n\r\nNote: Only the creator of an event or users with the EventsBotAdmin role assigned can cancel an event.", message)
		message = fmt.Sprintf("%s```", message)
	case "signup":
		message = fmt.Sprintf("%s\r\nHere's how to sign up to an event:", message)
		message = fmt.Sprintf("%s\r\n```%ssignup EventID @Username\r\n", message, config.CommandPrefix)
		message = fmt.Sprintf("%s\r\n    EventID: That weird looking 7 character identifier that uniquely identifies the event. These values are case sensitive so do take care to get it right. It's your key to participation, enjoyment and a deeper level of zen.", message)
		message = fmt.Sprintf("%s\r\n  @Username: The Discord user whom you wish to sign up to the event. Only the event creator and users with the EventsBotAdmin role assigned are allowed to sign users other than themselves up to an event.", message)
		message = fmt.Sprintf("%s\r\n\r\nNote: You can still sign up to an event even if it is already full. You will then be registered as a reserve for the event and promoted if someone leaves the event.", message)
		message = fmt.Sprintf("%s```", message)
	case "leave":
		message = fmt.Sprintf("%s\r\nHere's how to leave an event:", message)
		message = fmt.Sprintf("%s\r\n```%sleave EventID\r\n", message, config.CommandPrefix)
		message = fmt.Sprintf("%s\r\n    EventID: That weird looking 7 character identifier that uniquely identifies the event. These values are case sensitive so do take care to get it right. It's your key to participation, enjoyment and a deeper level of zen.", message)
		message = fmt.Sprintf("%s\r\n  @Username: The Discord user whom you wish to remove from the event. Only the event creator and users with the EventsBotAdmin role assigned are allowed to remove users other than themselves from an event.", message)
		message = fmt.Sprintf("%s```", message)
	case "impersonate":
		message = fmt.Sprintf("%s\r\nHere's how to impersonate a user:", message)
		message = fmt.Sprintf("%s\r\n```%simpersonate @Username\r\n", message, config.CommandPrefix)
		message = fmt.Sprintf("%s\r\n  @Username: The Discord user you wish to impersonate", message)
		message = fmt.Sprintf("%s\r\n\r\nNote: This will have the effect of any further commands you issue, until you've issued ~unimpersonate, behaving as if they originated from the specified user. This is dangerous of course and so only users with the EventsBotAdmin role assigned are allowed to issue this command. You have been warned.", message)
		message = fmt.Sprintf("%s```", message)
	case "unimpersonate":
		message = fmt.Sprintf("%s\r\nHere's how to stop impersonating a user:", message)
		message = fmt.Sprintf("%s\r\n```%sunimpersonate\r\n", message, config.CommandPrefix)
		message = fmt.Sprintf("%sYes, it's that simple", message)
		message = fmt.Sprintf("%s```", message)
	default:
		message = fmt.Sprintf("%s\r\nWait! What? Are you having me on? I don't know anything about %s", message, command[1])
		message = fmt.Sprintf("%s\r\nEventsBot is not amused :expressionless:", message)
	}

	s.ChannelMessageSend(m.ChannelID, message)
}

func Impersonate(s *discordgo.Session, m *discordgo.MessageCreate, command []string) {
	if !HasRole(s, m, "EventsBotAdmin") {
		message := fmt.Sprintf("Yo yo yo. Back up a second dude. You don't have permissions to impersonate other users.\r\nEventsBot will not stand for this :point_up:")
		message = fmt.Sprintf("%s\r\nFor help with impersonating users, type the following:\r\n```%shelp impersonate```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	user := ClanUser{
		UserName: command[1],
		Mention:  command[2],
	}

	impersonated = user
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s is now impersonated\r\nEventsBot is regarding this with some sense of apprehension :bust_in_silhouette:", impersonated.UserName))
}

func Unimpersonate(s *discordgo.Session, m *discordgo.MessageCreate, command []string) {
	impersonated = ClanUser{}
	s.ChannelMessageSend(m.ChannelID, "No more of this impersonation business!")
}

func Test(s *discordgo.Session, m *discordgo.MessageCreate, command []string) {

}

func IsUser(arg string) bool {
	return strings.HasPrefix(arg, "<@")
}

func IsDate(arg string) bool {
	_, err := time.Parse("2006-01-02", arg)
	return err == nil
}

func GetGuild(s *discordgo.Session, m *discordgo.MessageCreate) *discordgo.Guild {
	// Attempt to get the channel from the state.
	// If there is an error, fall back to the restapi
	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		channel, err = s.Channel(m.ChannelID)
		if err != nil {
			return nil
		}
	}

	// Attempt to get the guild from the state,
	// If there is an error, fall back to the restapi.
	guild, err := s.State.Guild(channel.GuildID)
	if err != nil {
		guild, err = s.Guild(channel.GuildID)
		if err != nil {
			return nil
		}
	}

	return guild
}

func HasRole(s *discordgo.Session, m *discordgo.MessageCreate, role string) bool {
	roleId := ""
	guild := GetGuild(s, m)
	for _, guildRole := range guild.Roles {
		if guildRole.Name == role {
			roleId = guildRole.ID
		}
	}
	found := false
	for _, member := range guild.Members {
		if member.User.Username == m.Author.Username {
			for _, memberrole := range member.Roles {
				if memberrole == roleId {
					found = true
				}
			}
		}
	}

	return found
}
