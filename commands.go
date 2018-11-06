package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// ListEvents is used to list all upcoming events on a specified (optional) date, for a specified (optional) user
// ~listevents
// ~listevents @username
// ~listevents date
// ~listevents date @username
func ListEvents(g *discordgo.Guild, s *discordgo.Session, m *discordgo.MessageCreate, command []string) {
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
		if isUser(command[1]) {
			listuser = m.Mentions[0].Username
		} else if isDate(command[1]) {
			specdate, _ = time.ParseInLocation("02/01/2006", command[1], defaultLocation)
		} else {
			message = fmt.Sprintf("Whoah, not so sure about those arguments. EventsBot is confused :confounded:")
			message = fmt.Sprintf("%s\r\nFor help with listing events, type the following:\r\n```%shelp listevents```", message, config.CommandPrefix)
			s.ChannelMessageSend(m.ChannelID, message)
			return
		}
	}

	// Check second argument
	if len(command) > 2 {
		if isUser(command[2]) {
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
		"$gte": time.Now().Add(-1 * time.Hour),
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

	c := mongoSession.DB(fmt.Sprintf("ClanEvents%s", g.ID)).C("Events")

	var results []ClanEvent
	err := c.Find(filter).Sort("dateTime").All(&results)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, ":scream::scream::scream:Something very weird happened when trying to read the events. Sorry but EventsBot has no answers for you :cry:")
		return
	}

	// Get all time zones
	ctz := mongoSession.DB(fmt.Sprintf("ClanEvents%s", g.ID)).C("TimeZones")
	var tzs []TimeZone
	tzLookup := make(map[string]TimeZone)
	err = ctz.Find(bson.M{}).Sort("abbrev").All(&tzs)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, ":scream::scream::scream:Something very weird happened when trying to read the events. Sorry but EventsBot has no answers for you :cry:")
		return
	}
	for _, tz := range tzs {
		tzLookup[tz.Abbrev] = tz
	}

	var reply string
	if specdate.IsZero() {
		reply = fmt.Sprintf("%s - Upcoming events", g.Name)
	} else {
		reply = fmt.Sprintf("%s - Events on %s", g.Name, specdate.Format("Mon 02/01/2006"))
	}
	reply = fmt.Sprintf("%s for %s\r\n", reply, listuser)
	if len(results) == 0 {
		reply = fmt.Sprintf("%sZip. Nothing. Nada.\r\nWhat nonsense is this? EventsBot does not approve :frowning2:", reply)
	} else {
		reply = fmt.Sprintf("%s```", reply)
		for _, event := range results {
			tzInfo := ""
			eventLocation := defaultLocation
			if event.TimeZone != "" {
				tzInfo = fmt.Sprintf(" (%s)", event.TimeZone)
				eventLocation, _ = time.LoadLocation(tzLookup[event.TimeZone].Location)
			}
			freeSpace := event.TeamSize - len(event.Participants)
			//reply = fmt.Sprintf("%s%8v: %s%s - %s", reply, event.EventID, event.DateTime.In(defaultLocation).Format("Mon 02/01 15:04"), tzInfo, event.Name)
			reply = fmt.Sprintf("%s%8v: %s%s - %s", reply, event.EventID, event.DateTime.In(eventLocation).Format("Mon 02/01 15:04"), tzInfo, event.Name)

			// Add players to message
			if len(event.Participants) > 0 {
				reply = fmt.Sprintf("%s\r\n Players:", reply)

				for _, participant := range event.Participants {
					reply = fmt.Sprintf("%s %s,", reply, participant.DisplayName())
				}
				// Remove trailing comma
				reply = fmt.Sprintf("%s", strings.TrimSuffix(reply, ","))
			}

			// Add reserves to message
			if len(event.Reserves) > 0 {
				reply = fmt.Sprintf("%s\r\nReserves:", reply)

				for _, reserve := range event.Reserves {
					reply = fmt.Sprintf("%s %s,", reply, reserve.DisplayName())
				}
				// Remove trailing comma
				reply = fmt.Sprintf("%s", strings.TrimSuffix(reply, ","))
			}

			// Add status to message
			reply = fmt.Sprintf("%s\r\n  Status: ", reply)
			switch freeSpace {
			case 0:
				reply = fmt.Sprintf("%sFULL", reply)
			case 1:
				reply = fmt.Sprintf("%s1 Space", reply)
			default:
				reply = fmt.Sprintf("%s%d Spaces", reply, freeSpace)
			}

			reply = fmt.Sprintf("%s\r\n----------------------------------------\r\n", reply)
		}
		reply = fmt.Sprintf("%s```", reply)
	}

	s.ChannelMessageSend(m.ChannelID, reply)
}

// Details is used to display detailed information on a specified event
// ~details EventID
func Details(g *discordgo.Guild, s *discordgo.Session, m *discordgo.MessageCreate, command []string) {
	message := ""

	// Test for correct number of arguments
	if len(command) != 2 {
		message = fmt.Sprintf("Whoah, not so sure about those arguments. EventsBot is confused :thinking:")
		message = fmt.Sprintf("%s\r\nFor help with getting the details of an event, type the following:\r\n```%shelp details```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	// Find event in DB
	c := mongoSession.DB(fmt.Sprintf("ClanEvents%s", g.ID)).C("Events")

	var event ClanEvent
	err := c.Find(bson.M{"eventId": command[1]}).One(&event)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("EventsBot could find no such event. Are you sure you got that Event ID of %s right? Them's finicky numbers. Remember, they're case sensitive :grimacing:", command[1]))
		return
	}

	// Get time zone
	tzInfo := ""
	eventLocation := defaultLocation
	ctz := mongoSession.DB(fmt.Sprintf("ClanEvents%s", g.ID)).C("TimeZones")
	if event.TimeZone != "" {
		var tz TimeZone
		err = ctz.Find(bson.M{"abbrev": event.TimeZone}).One(&tz)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("EventsBot had trouble interpreting the time zone information of this event. Are we anywhere near a worm hole perhaps? :no_mouth:"))
			return
		}
		tzInfo = tz.Abbrev
		eventLocation, _ = time.LoadLocation(tz.Location)
	}

	message = fmt.Sprintf("**EventID:** %s", event.EventID)
	message = fmt.Sprintf("%s\r\n**Creator:** %s", message, event.Creator.Mention())
	message = fmt.Sprintf("%s\r\n**Date:** %s", message, event.DateTime.In(eventLocation).Format("Mon 02/01/2006"))
	message = fmt.Sprintf("%s\r\n**Time:** %s for %d hours", message, event.DateTime.In(eventLocation).Format("15:04"), event.Duration)
	if event.TimeZone != "" {
		message = fmt.Sprintf("%s\r\n**Time Zone:** %s", message, tzInfo)
	}
	message = fmt.Sprintf("%s\r\n**Name:** %s", message, event.Name)
	message = fmt.Sprintf("%s\r\n**Description:** %s", message, event.Description)
	message = fmt.Sprintf("%s\r\n**Team Size:** %d of %d", message, len(event.Participants), event.TeamSize)
	if len(event.Participants) > 0 {
		message = fmt.Sprintf("%s\r\n**Participants:**", message)
		for _, participant := range event.Participants {
			message = fmt.Sprintf("%s\r\n   -  %s", message, participant.Mention())
		}
	}
	if len(event.Reserves) > 0 {
		message = fmt.Sprintf("%s\r\n**Reserves:**", message)
		for _, reserve := range event.Reserves {
			message = fmt.Sprintf("%s\r\n    - %s", message, reserve.Mention())
		}
	}

	s.ChannelMessageSend(m.ChannelID, message)
}

// NewEvent is used to create a new event
// ~newevent Date Time (TimeZone) Duration Name Description Size
func NewEvent(g *discordgo.Guild, s *discordgo.Session, m *discordgo.MessageCreate, command []string) {
	message := ""

	// Test for correct number of arguments
	if len(command) < 7 || len(command) > 8 {
		message = fmt.Sprintf("Whoah, not so sure about those arguments. EventsBot is confused :thinking:")
		message = fmt.Sprintf("%s\r\nFor help with creating a new event, type the following:\r\n```%shelp newevent```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	tzOffset := len(command) - 7 // offset value to use later for all parameters that come after the timezone (if present)
	locAbbr := ""
	var usrLocation *time.Location
	// Test for time zone specification
	if tzOffset > 0 {
		if !strings.HasPrefix(command[2+tzOffset], "(") || !strings.HasSuffix(command[2+tzOffset], ")") {
			message = fmt.Sprintf("Is %s supposed to be a time zone? Please put time zones in brackets :point_up:", command[2+tzOffset])
			message = fmt.Sprintf("%s\r\nFor help with creating a new event, type the following:\r\n```%shelp newevent```", message, config.CommandPrefix)
			s.ChannelMessageSend(m.ChannelID, message)
			return
		}

		locAbbr = command[2+tzOffset]
		locAbbr = strings.TrimPrefix(locAbbr, "(")
		locAbbr = strings.TrimSuffix(locAbbr, ")")

		// Check if timezone is known
		ctz := mongoSession.DB(fmt.Sprintf("ClanEvents%s", g.ID)).C("TimeZones")
		var tz TimeZone
		err := ctz.Find(bson.M{"abbrev": locAbbr}).One(&tz)
		if err != nil || tz.Location == "" {
			message = fmt.Sprintf("EventsBot doesn't know that %s time zone. Can we stick to time zones on earth please?", command[2+tzOffset])
			message = fmt.Sprintf("%s\r\nTo see a list of available time zones, type the following:\r\n```%listtimezones```", message, config.CommandPrefix)
			s.ChannelMessageSend(m.ChannelID, message)
			return
		}

		usrLocation, err = time.LoadLocation(tz.Location)
		if err != nil {
			message = fmt.Sprintf("EventsBot is having trouble working with this time zone. Are we anywhere near a worm hole perhaps? :no_mouth:")
			message = fmt.Sprintf("%s\r\nFor help with creating a new event, type the following:\r\n```%shelp newevent```", message, config.CommandPrefix)
			s.ChannelMessageSend(m.ChannelID, message)
			return
		}
	} else {
		usrLocation, locAbbr = getLocation(g, s, m)
	}

	// Test for date and time arguments
	datetime := fmt.Sprintf("%s %s", command[1], command[2])
	dt, err := time.ParseInLocation("02/01/2006 15:04", datetime, usrLocation)
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
	duration, err := strconv.Atoi(command[3+tzOffset])
	if err != nil {
		message = fmt.Sprintf("What kind of a duration is %s? EventsBot needs a vacation of %s weeks :beach:", command[3+tzOffset], command[3+tzOffset])
		message = fmt.Sprintf("%s\r\nFor help with creating a new event, type the following:\r\n```%shelp newevent```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	// Test for name
	if len(command[4+tzOffset]) > 50 {
		message = fmt.Sprintf("That's a very long name right there. You realise EventsBot has to memorise these things? Have a heart and keep it under 50 characters please. :triumph:")
		message = fmt.Sprintf("%s\r\nFor help with creating a new event, type the following:\r\n```%shelp newevent```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	// Test for description
	if len(command[5+tzOffset]) > 150 {
		message = fmt.Sprintf("That's a very long description right there. You realise EventsBot has to memorise these things? Have a heart and keep it under 150 characters please. :triumph:")
		message = fmt.Sprintf("%s\r\nFor help with creating a new event, type the following:\r\n```%shelp newevent```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	// Test for size
	teamSize, err := strconv.Atoi(command[6+tzOffset])
	if err != nil {
		message = fmt.Sprintf("How many players you say? %s? EventsBot wouldn't do that if he were you :speak_no_evil:", command[6+tzOffset])
		message = fmt.Sprintf("%s\r\nFor help with creating a new event, type the following:\r\n```%shelp newevent```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	//newid, _ := baseconv.Encode62FromDec(time.Now().Format("050415020106")) // Convert the current DateTime in reverse order of significance (ssmmHHDDMMYY) to Base62
	newid := getEventID(time.Now())
	curUser := impersonated
	if curUser.UserName == "" {
		curUser = ClanUser{
			UserName: m.Author.Username,
			UserID:   m.Author.ID,
			Nickname: getNickname(g, s, m.Author.ID),
			DateTime: time.Now(),
		}
	}
	newEvent := ClanEvent{
		EventID:     newid,
		Creator:     curUser,
		DateTime:    dt,
		TimeZone:    locAbbr,
		Duration:    duration,
		Name:        command[4+tzOffset],
		Description: command[5+tzOffset],
		TeamSize:    teamSize,
		Full:        false,
	}

	c := mongoSession.DB(fmt.Sprintf("ClanEvents%s", g.ID)).C("Events")
	err = c.Insert(newEvent)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, ":scream::scream::scream:Something very weird happened when trying to create this event. Sorry but EventsBot has no answers for you :cry:")
		return
	}

	message = fmt.Sprintf("Woohoo! A new event has been created by %s. EventsBot is most pleased :ok_hand:", newEvent.Creator.Mention())
	message = fmt.Sprintf("%s\r\nEvent ID: **%s**", message, newEvent.EventID)
	message = fmt.Sprintf("%s\r\n\r\nTo sign up for this event, type the following:", message)
	message = fmt.Sprintf("%s\r\n```%ssignup %s```", message, config.CommandPrefix, newEvent.EventID)
	s.ChannelMessageSend(m.ChannelID, message)

	signupCmd := []string{"signup", newEvent.EventID}
	Signup(g, s, m, signupCmd)
}

// CancelEvent is used to delete a specified event
// ~cancelevent EventID
func CancelEvent(g *discordgo.Guild, s *discordgo.Session, m *discordgo.MessageCreate, command []string) {
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
			UserID:   m.Author.ID,
			Nickname: getNickname(g, s, m.Author.ID),
			DateTime: time.Now(),
		}
	}

	// Find event in DB
	c := mongoSession.DB(fmt.Sprintf("ClanEvents%s", g.ID)).C("Events")

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
	} else if hasRole(g, s, m, "EventsBotAdmin") {
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

	message = fmt.Sprintf("Tragedy! %s's event, %s, has been cancelled.", event.Creator.Mention(), event.Name)
	message = fmt.Sprintf("%s\r\n\r\n\"We don't have commit. Repeat. We are decommissioning the committal of the launch. It is now a negatory launch phase. We are in a no fly, no go phase. That is a November Gorgon phase, of non-flying. And we're gonna say 'goodnight, thank you, good work, over and out'\".", message)
	message = fmt.Sprintf("%s\r\n\r\nEventsBot will cry himself to sleep tonight :sob:", message)
	s.ChannelMessageSend(m.ChannelID, message)
}

// Signup is used to sign the author or a specified user up to an event
// ~signup EventID
// ~signup EventID @Username
func Signup(g *discordgo.Guild, s *discordgo.Session, m *discordgo.MessageCreate, command []string) {
	message := ""

	curUser := impersonated
	curUser.DateTime = time.Now()
	if curUser.UserName == "" {
		curUser = ClanUser{
			UserName: m.Author.Username,
			UserID:   m.Author.ID,
			Nickname: getNickname(g, s, m.Author.ID),
			DateTime: time.Now(),
		}
	}

	signupUsers := []ClanUser{}
	
	if len(command) < 2 {
		message = fmt.Sprintf("Come on! Surely you're not expecting me to guess which event you're trying to sign up to :confused:")
		message = fmt.Sprintf("%s\r\nFor help with signing up to events, type the following:\r\n```%shelp signup```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	// Check first argument
	if len(command) > 2 {
		for i := 2; i < len(command); i++ {
			if isUser(command[i]) {
				// Find user in list of mentions
				for _, mentionedUser := range m.Mentions {
					if strings.Replace(command[i], "!", "", 1) == mentionedUser.Mention() {
						signupUser := ClanUser{
							UserName: mentionedUser.Username,
							UserID:   mentionedUser.ID,
							Nickname: getNickname(g, s, mentionedUser.ID),
							DateTime: time.Now(),
						}
						signupUsers = append(signupUsers, signupUser)
					}
				}
			} else {
				message = fmt.Sprintf("Whoah, not so sure about those arguments. EventsBot is confused :confounded:")
				message = fmt.Sprintf("%s\r\n%s doesn't look like anyone I recognise.", message, command[i])
				message = fmt.Sprintf("%s\r\nFor help with signing up to events, type the following:\r\n```%shelp signup```", message, config.CommandPrefix)
				s.ChannelMessageSend(m.ChannelID, message)
				return
			}
		}
	} else {
		signupUsers = append(signupUsers, curUser)
	}

	// Find event in DB
	c := mongoSession.DB(fmt.Sprintf("ClanEvents%s", g.ID)).C("Events")

	var event ClanEvent
	err := c.Find(bson.M{"eventId": command[1]}).One(&event)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("EventsBot could find no such event. Are you sure you got that Event ID of %s right? Them's finicky numbers. Remember, they're case sensitive :grimacing:", command[1]))
		return
	}

	// If different user is specified, check that current user has permissions
	if signupUsers[0] != curUser || len(signupUsers) > 1 {
		allowed := false
		if event.Creator.UserName == curUser.UserName {
			allowed = true
		} else if hasRole(g, s, m, "EventsBotAdmin") {
			allowed = true
		}

		if !allowed {
			message = fmt.Sprintf("Yo yo yo. Back up a second dude. You don't have permissions to sign other users up to events.\r\nEventsBot will not stand for this :point_up:")
			message = fmt.Sprintf("%s\r\nFor help with signing up to events, type the following:\r\n```%shelp signup```", message, config.CommandPrefix)
			s.ChannelMessageSend(m.ChannelID, message)
			return
		}
	}

	// Check if any of the specified users are already signed up for this event
	for i1, signupUser := range signupUsers {
		for _, participant := range event.Participants {
			if participant.UserName == signupUser.UserName {
				if signupUser.UserName == curUser.UserName {
					s.ChannelMessageSend(m.ChannelID, "You are already signed up to this event.\r\nEventsBot hasn't got time for your shenanigans :rolling_eyes:")
				} else {
					s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s is already signed up to this event.\r\nEventsBot hasn't got time for your shenanigans :rolling_eyes:", signupUser.DisplayName()))
				}
				return
			}
		}
		for _, reserve := range event.Reserves {
			if reserve.UserName == signupUser.UserName {
				if signupUser.UserName == curUser.UserName {
					s.ChannelMessageSend(m.ChannelID, "You are already a reserve for this event.\r\nCan you just relax please? EventsBot will let you know if a space opens up. Don't call us, we'll call you. :rolling_eyes:")
				} else {
					s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s is already a reserve for this event.\r\nCan you just relax please? EventsBot will let you know if a space opens up. Don't call us, we'll call you. :rolling_eyes:", signupUser.DisplayName()))
				}
				return
			}
		}
		for i2 := 20; i2 < i1; i2++ {
			if signupUsers[i1].UserName == signupUsers[i2].UserName {
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("How many times do you want to sign %s up for this event.\r\nNo one is _that_ important. :confused:", signupUser.DisplayName()))
				return
			}
		}
	}

	// Sign up all specified users
	for _, signupUser := range signupUsers {
		// Check if event is full
		if len(event.Participants) >= event.TeamSize {
			s.ChannelMessageSend(m.ChannelID, "Oh noes! This event is already full :cry:\r\nBut don't worry, EventsBot will put you on the reserves list and notify you if someone leaves.")
			for _, reserve := range event.Reserves {
				if reserve.UserName == curUser.UserName {
					continue
				}
			}
			event.Reserves = append(event.Reserves, signupUser)
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s is now signed up as a reserve for %s's event, %s.\r\nEventsBot approves :thumbsup:", signupUser.Mention(), event.Creator.Mention(), event.Name))
		} else {
			event.Participants = append(event.Participants, signupUser)
			event.Full = len(event.Participants) >= event.TeamSize
			message := fmt.Sprintf("%s is now signed up for %s's event, %s.\r\n", signupUser.Mention(), event.Creator.Mention(), event.Name)
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
	err = c.Update(bson.M{"eventId": command[1]}, event)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, ":scream::scream::scream:Something very weird happened when trying to update the event. Sorry but EventsBot has no answers for you :cry:")
		return
	}
}

// Leave is used to remove the author or specified user from an event
// ~leave EventID
// ~leave EventID @Username
func Leave(g *discordgo.Guild, s *discordgo.Session, m *discordgo.MessageCreate, command []string) {
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
			UserID:   m.Author.ID,
			Nickname: getNickname(g, s, m.Author.ID),
			DateTime: time.Now(),
		}
	}

	removeUser := curUser

	// Check first argument
	if len(command) > 2 {
		if isUser(command[2]) {
			removeUser = ClanUser{
				UserName: m.Mentions[0].Username,
				UserID:   m.Mentions[0].ID,
				Nickname: getNickname(g, s, m.Mentions[0].ID),
				DateTime: time.Now(),
			}
		} else {
			message = fmt.Sprintf("Whoah, not so sure about those arguments. EventsBot is confused :confounded:")
			message = fmt.Sprintf("%s\r\nFor help with signing up to events, type the following:\r\n```%shelp signup```", message, config.CommandPrefix)
			s.ChannelMessageSend(m.ChannelID, message)
			return
		}
	}

	// Find event in DB
	c := mongoSession.DB(fmt.Sprintf("ClanEvents%s", g.ID)).C("Events")

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
		} else if hasRole(g, s, m, "EventsBotAdmin") {
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
	participantIndex := -1
	for i, participant := range event.Participants {
		if participant.UserName == removeUser.UserName {
			participantIndex = i
		}
	}

	if participantIndex != -1 {
		// Remove participant from event
		event.Participants = append(event.Participants[:participantIndex], event.Participants[participantIndex+1:]...)
		message = fmt.Sprintf("Well okay then, %s has been removed from %s's event, %s\r\nEventsBot is sad to see you go :disappointed_relieved:", removeUser.Mention(), event.Creator.Mention(), event.Name)

		// Move first reserve into participants
		if len(event.Reserves) > 0 {
			message = fmt.Sprintf("%s\r\nBut hey! %s is on reserve so we're golden.\r\nEventsBot is relieved :relieved:", message, event.Reserves[0].Mention())
			reserve := ClanUser{
				UserName: event.Reserves[0].UserName,
				UserID:   event.Reserves[0].UserID,
				Nickname: event.Reserves[0].Nickname,
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
			if event.TeamSize-len(event.Participants) == 1 {
				message = fmt.Sprintf("%s\r\nThere is now one space left\r\n", message)
			} else {
				message = fmt.Sprintf("%s\r\nThere are now %d spaces left\r\n", message, event.TeamSize-len(event.Participants))
			}
		}
	}

	// Check if user is a reserve for this event
	reserveIndex := -1
	for i, reserve := range event.Reserves {
		if reserve.UserName == removeUser.UserName {
			reserveIndex = i
		}
	}

	if reserveIndex != -1 {
		// Remove reserve from event
		event.Reserves = append(event.Reserves[:reserveIndex], event.Reserves[reserveIndex+1:]...)
		message = fmt.Sprintf("Well okay then, %s has been removed as a reserve from %s's event, %s\r\nEventsBot is sad to see you go :disappointed_relieved:", removeUser.Mention(), event.Creator.Mention(), event.Name)
	}

	if participantIndex == -1 && reserveIndex == -1 {
		if curUser.UserName == removeUser.UserName {
			s.ChannelMessageSend(m.ChannelID, "You are not signed up to this event.\r\nEventsBot does not find your jokes particularly funny :rolling_eyes:")
		} else {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s is not signed up to this event.\r\nEventsBot does not find your jokes particularly funny :rolling_eyes:", removeUser.DisplayName()))
		}
		return
	}

	err = c.Update(bson.M{"eventId": command[1]}, event)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, ":scream::scream::scream:Something very weird happened when trying to update the event. Sorry but EventsBot has no answers for you :cry:")
		return
	}
	s.ChannelMessageSend(m.ChannelID, message)
}

// BotHelp is used to display a list of available commands or instructions on using a specified command
func BotHelp(g *discordgo.Guild, s *discordgo.Session, m *discordgo.MessageCreate, command []string) {
	if len(command) == 1 {
		command = append(command, "nothing")
	}

	message := fmt.Sprintf("Need some help with the __%s__ command? EventsBot is happy to oblige :nerd:", command[1])

	switch command[1] {
	case "nothing":
		message = "```List of EventsBot commands:"
		message = fmt.Sprintf("%s\r\n    %slistevents", message, config.CommandPrefix)
		message = fmt.Sprintf("%s\r\n    %sdetails", message, config.CommandPrefix)
		message = fmt.Sprintf("%s\r\n    %snewevent", message, config.CommandPrefix)
		message = fmt.Sprintf("%s\r\n    %scancelevent", message, config.CommandPrefix)
		message = fmt.Sprintf("%s\r\n    %ssignup", message, config.CommandPrefix)
		message = fmt.Sprintf("%s\r\n    %sleave", message, config.CommandPrefix)
		message = fmt.Sprintf("%s\r\n    %swisdom", message, config.CommandPrefix)
		message = fmt.Sprintf("%s\r\n    %slisttimezones", message, config.CommandPrefix)
		message = fmt.Sprintf("%s\r\n\r\nCommands that require Admin priviledges", message)
		message = fmt.Sprintf("%s\r\n    %saddserver", message, config.CommandPrefix)
		message = fmt.Sprintf("%s\r\n    %saddtimezone", message, config.CommandPrefix)
		message = fmt.Sprintf("%s\r\n    %sremovetimezone", message, config.CommandPrefix)
		message = fmt.Sprintf("%s\r\n    %sroletimezone", message, config.CommandPrefix)
		//message = fmt.Sprintf("%s\r\n    %simpersonate", message, config.CommandPrefix)
		//message = fmt.Sprintf("%s\r\n    %sunimpersonate", message, config.CommandPrefix)
		message = fmt.Sprintf("%s```", message)
		message = fmt.Sprintf("%sYou can get help on any of these commands by typing %shelp followed by the name of the command", message, config.CommandPrefix)
	case "listevents":
		message = fmt.Sprintf("%s\r\nHere's how to get a list of upcoming events:", message)
		message = fmt.Sprintf("%s\r\n```%slistevents [Date] [@Username]\r\n", message, config.CommandPrefix)
		message = fmt.Sprintf("%s\r\n       Date: The date for which you want to see events. This value is optional", message)
		message = fmt.Sprintf("%s\r\n  @Username: The Discord user for which you want to see events. This value is optional.", message)
		message = fmt.Sprintf("%s\r\n\r\nNote: Both the date and username values are optional. You can specify either, neither or both but then they must be in the order shown above. If you omit the date, you will be shown all upcoming events and if you omit the user you will be shown events for all users.", message)
		message = fmt.Sprintf("%s```", message)
	case "details":
		message = fmt.Sprintf("%s\r\nHere's how to get details for an event:", message)
		message = fmt.Sprintf("%s\r\n```%sdetails EventID\r\n", message, config.CommandPrefix)
		message = fmt.Sprintf("%s\r\n    EventID: That weird looking 7 character identifier that uniquely identifies the event. These values are case sensitive so do take care to get it right. It's your key to participation, enjoyment and a deeper level of zen.", message)
		message = fmt.Sprintf("%s```", message)
	case "newevent":
		message = fmt.Sprintf("%s\r\nHere's how to create a new event:", message)
		message = fmt.Sprintf("%s\r\n```%snewevent Date Time (TimeZone) Duration Name Description Size\r\n", message, config.CommandPrefix)
		message = fmt.Sprintf("%s\r\n       Date: In the format DD/MM/YYYY", message)
		message = fmt.Sprintf("%s\r\n       Time: In the format HH:MM (24 hour clock)", message)
		message = fmt.Sprintf("%s\r\n   TimeZone: A time zone abbreviation between brackets", message)
		message = fmt.Sprintf("%s\r\n   Duration: Number of hours the event will last", message)
		message = fmt.Sprintf("%s\r\n       Name: A name for your event. Surround it in quotes if it's more than one word", message)
		message = fmt.Sprintf("%s\r\nDescription: A longer description of your event. You totally want to surround this one in quotes", message)
		message = fmt.Sprintf("%s\r\n   TeamSize: Just a number denoting how many players can sign up", message)
		message = fmt.Sprintf("%s\r\n\r\nNote: Specifying a time zone is optional, as can be seen in the example below. If no time zone is specified, the role default time zone will be used.", message)
		message = fmt.Sprintf("%s```", message)
		message = fmt.Sprintf("%s\r\n\r\nHere's an example for you:", message)
		message = fmt.Sprintf("%s\r\n```%snewevent %s 20:00 2 \"Normal Leviathan\" \"Fresh start of Leviathan raid\" 6```", message, config.CommandPrefix, time.Now().Format("02/01/2006"))
		message = fmt.Sprintf("%s\r\nThis will create a 2 hour event to start at 8pm tonight and which will allow 6 people to sign up", message)
	case "cancelevent":
		message = fmt.Sprintf("%s\r\nHere's how to cancel an event:", message)
		message = fmt.Sprintf("%s\r\n```%scancelevent EventID\r\n", message, config.CommandPrefix)
		message = fmt.Sprintf("%s\r\n    EventID: That weird looking 7 character identifier that uniquely identifies the event. These values are case sensitive so do take care to get it right. It's your key to participation, enjoyment and a deeper level of zen.", message)
		message = fmt.Sprintf("%s\r\n\r\nNote: Only the creator of an event or users with the EventsBotAdmin role assigned can cancel an event.", message)
		message = fmt.Sprintf("%s```", message)
	case "signup":
		message = fmt.Sprintf("%s\r\nHere's how to sign up to an event:", message)
		message = fmt.Sprintf("%s\r\n```%ssignup EventID [@Username] [@Username] ...\r\n", message, config.CommandPrefix)
		message = fmt.Sprintf("%s\r\n    EventID: That weird looking 7 character identifier that uniquely identifies the event. These values are case sensitive so do take care to get it right. It's your key to participation, enjoyment and a deeper level of zen.", message)
		message = fmt.Sprintf("%s\r\n  @Username: List of Discord users whom you wish to sign up to the event. Only the event creator and users with the EventsBotAdmin role assigned are allowed to sign users other than themselves up to an event. This value is optional.", message)
		message = fmt.Sprintf("%s\r\n\r\nNote: You can still sign up to an event even if it is already full. You will then be registered as a reserve for the event and promoted if someone leaves the event.", message)
		message = fmt.Sprintf("%s```", message)
	case "leave":
		message = fmt.Sprintf("%s\r\nHere's how to leave an event:", message)
		message = fmt.Sprintf("%s\r\n```%sleave EventID [@Username]\r\n", message, config.CommandPrefix)
		message = fmt.Sprintf("%s\r\n    EventID: That weird looking 7 character identifier that uniquely identifies the event. These values are case sensitive so do take care to get it right. It's your key to participation, enjoyment and a deeper level of zen.", message)
		message = fmt.Sprintf("%s\r\n  @Username: The Discord user whom you wish to remove from the event. Only the event creator and users with the EventsBotAdmin role assigned are allowed to remove users other than themselves from an event. This value is optional.", message)
		message = fmt.Sprintf("%s```", message)
	case "impersonate":
		message = fmt.Sprintf("%s\r\nHere's how to impersonate a user:", message)
		message = fmt.Sprintf("%s\r\n```%simpersonate @Username\r\n", message, config.CommandPrefix)
		message = fmt.Sprintf("%s\r\n  @Username: The Discord user you wish to impersonate", message)
		message = fmt.Sprintf("%s\r\n\r\nNote: This will have the effect of any further commands you issue, until you've issued %sunimpersonate, behaving as if they originated from the specified user. This is dangerous of course and so only users with the EventsBotAdmin role assigned are allowed to issue this command. You have been warned.", message, config.CommandPrefix)
		message = fmt.Sprintf("%s```", message)
	case "unimpersonate":
		message = fmt.Sprintf("%s\r\nHere's how to stop impersonating a user:", message)
		message = fmt.Sprintf("%s\r\n```%sunimpersonate\r\n", message, config.CommandPrefix)
		message = fmt.Sprintf("%sYes, it's that simple", message)
		message = fmt.Sprintf("%s```", message)
	case "wisdom":
		message = fmt.Sprintf("%s\r\nHere's how to obtain a nugget of wisdom:", message)
		message = fmt.Sprintf("%s\r\n```%swisdom\r\n", message, config.CommandPrefix)
		message = fmt.Sprintf("%sYes, it's that simple. Just ask and you shall receive.", message)
		message = fmt.Sprintf("%s```", message)
	case "addnaughtylist":
		message = fmt.Sprintf("%s\r\nHere's how to add a user to the naughty list:", message)
		message = fmt.Sprintf("%s\r\n```%saddnaughtylist @Username\r\n", message, config.CommandPrefix)
		message = fmt.Sprintf("%s\r\n  @Username: The Discord user you wish to add to the naughty list", message)
		message = fmt.Sprintf("%s```", message)
	case "addserver":
		message = fmt.Sprintf("%s\r\nHere's how to add a server to EventsBot:", message)
		message = fmt.Sprintf("%s\r\n```%saddserver\r\n", message, config.CommandPrefix)
		message = fmt.Sprintf("%sYes, it's that simple", message)
		message = fmt.Sprintf("%s```", message)
	case "addtimezone":
		message = fmt.Sprintf("%s\r\nHere's how to add a time zone to EventsBot:", message)
		message = fmt.Sprintf("%s\r\n```%saddtimezone Abbrev Location\r\n", message, config.CommandPrefix)
		message = fmt.Sprintf("%s\r\n     Abbrev: Abbreviation to be used for this time zone (ie. ET, CT, etc.)", message)
		message = fmt.Sprintf("%s\r\n   Location: A location that represents the time zone (conforms to the tz database naming convention)", message)
		message = fmt.Sprintf("%s\r\n\r\nNote: For more information on the tz database naming convention, see https://en.wikipedia.org/wiki/Tz_database", message)
		message = fmt.Sprintf("%s\r\n\r\nNote: EventsBot automatically adjusts times based on the specified location's Daylight Saving convention.", message)
		message = fmt.Sprintf("%s```", message)
	case "removetimezone":
		message = fmt.Sprintf("%s\r\nHere's how to remove a time zone from EventsBot:", message)
		message = fmt.Sprintf("%s\r\n     Abbrev: Abbreviation for the time zone to be removed", message)
		message = fmt.Sprintf("%s```", message)
	case "listtimezones":
		message = fmt.Sprintf("%s\r\nHere's how to obtain a list of time zones:", message)
		message = fmt.Sprintf("%s\r\n```%slisttimezones\r\n", message, config.CommandPrefix)
		message = fmt.Sprintf("%sYes, it's that simple. Just ask and you shall receive.", message)
		message = fmt.Sprintf("%s```", message)
	case "roletimezone":
		message = fmt.Sprintf("%s\r\nHere's how to associate a time zone to a server role:", message)
		message = fmt.Sprintf("%s\r\n     Abbrev: Abbreviation provided when '%saddtimezone' command was issued", message, config.CommandPrefix)
		message = fmt.Sprintf("%s```", message)
	default:
		message = fmt.Sprintf("%s\r\nWait! What? Are you having me on? I don't know anything about %s", message, command[1])
		message = fmt.Sprintf("%s\r\nEventsBot is not amused :expressionless:", message)
	}

	s.ChannelMessageSend(m.ChannelID, message)
}

// Impersonate is used to assume the identity of another Discord user and issue commands on that user's behalf
func Impersonate(g *discordgo.Guild, s *discordgo.Session, m *discordgo.MessageCreate, command []string) {
	message := ""

	if !hasRole(g, s, m, "EventsBotAdmin") {
		message = fmt.Sprintf("Yo yo yo. Back up a second dude. You don't have permissions to impersonate other users.\r\nEventsBot will not stand for this :point_up:")
		message = fmt.Sprintf("%s\r\nFor help with impersonating users, type the following:\r\n```%shelp impersonate```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	// Test for correct number of arguments
	if len(command) > 2 {
		message = fmt.Sprintf("Whoah, not so sure about those arguments. EventsBot is confused :thinking:")
		message = fmt.Sprintf("%s\r\nFor help with impersonating users, type the following:\r\n```%shelp impersonate```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	// Check first argument
	if !isUser(command[1]) {
		message = fmt.Sprintf("Whoah, not so sure about those arguments. EventsBot is confused :confounded:")
		message = fmt.Sprintf("%s\r\nFor help with impersonating users, type the following:\r\n```%shelp impersonate```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	user := ClanUser{
		UserName: m.Mentions[0].Username,
		UserID:   m.Mentions[0].ID,
		Nickname: getNickname(g, s, m.Mentions[0].ID),
		DateTime: time.Now(),
	}

	impersonated = user
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s is now impersonated\r\nEventsBot is regarding this with some sense of apprehension :bust_in_silhouette:", impersonated.DisplayName()))
}

// Unimpersonate is used to return to the original user's identity after impersonating another user
func Unimpersonate(g *discordgo.Guild, s *discordgo.Session, m *discordgo.MessageCreate, command []string) {
	impersonated = ClanUser{}
	s.ChannelMessageSend(m.ChannelID, "No more of this impersonation business!")
}

// Test is used to simply check that the bot is online and responding
func Test(g *discordgo.Guild, s *discordgo.Session, m *discordgo.MessageCreate, command []string) {
	t := time.Now()
	//newid, _ := baseconv.Encode36FromDec(time.Now().Format("05041502")) // Convert the current DateTime in reverse order of significance (ssmmHHDDMMYY) to Base62
	//newid = strings.ToUpper(newid)
	newid := getEventID(t)
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s", newid))

	fmt.Printf("%s - %s\r\n", t.Format("02 15 04 05"), getEventID(t))
	fmt.Printf("%s - %s\r\n", t.Add(1*time.Second).Format("02 15 04 05"), getEventID(t.Add(1*time.Second)))
	fmt.Printf("%s - %s\r\n", t.Add(2*time.Second).Format("02 15 04 05"), getEventID(t.Add(2*time.Second)))
	fmt.Printf("%s - %s\r\n", t.Add(3*time.Second).Format("02 15 04 05"), getEventID(t.Add(3*time.Second)))
	fmt.Printf("%s - %s\r\n", t.Add(4*time.Second).Format("02 15 04 05"), getEventID(t.Add(4*time.Second)))
	fmt.Printf("%s - %s\r\n", t.Add(5*time.Second).Format("02 15 04 05"), getEventID(t.Add(5*time.Second)))
	fmt.Printf("%s - %s\r\n", t.Add(6*time.Second).Format("02 15 04 05"), getEventID(t.Add(6*time.Second)))
	fmt.Printf("%s - %s\r\n", time.Date(2018, time.July, 01, 0, 0, 0, 0, time.UTC).Format("02 15 04 05"), getEventID(time.Date(2018, time.July, 01, 0, 0, 0, 0, time.UTC)))
	fmt.Printf("%s - %s\r\n", time.Date(2018, time.July, 31, 23, 59, 59, 0, time.UTC).Format("02 15 04 05"), getEventID(time.Date(2018, time.July, 31, 23, 59, 59, 0, time.UTC)))
}

// Wisdom is used to deliver a nugget of wisdom
func Wisdom(g *discordgo.Guild, s *discordgo.Session, m *discordgo.MessageCreate, command []string) {
	message := getInsult(m.Author.Mention())
	sendMessage(m.ChannelID, message)
}

// AddNaughty is used to add a user to the naughty list
func AddNaughty(g *discordgo.Guild, s *discordgo.Session, m *discordgo.MessageCreate, command []string) {
	message := ""

	// Test for correct number of arguments
	if len(command) != 2 {
		message = fmt.Sprintf("Whoah, not so sure about those arguments. EventsBot is confused :thinking:")
		message = fmt.Sprintf("%s\r\nFor help with adding a user to the naughty list, type the following:\r\n```%shelp addnaughtylist```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	var addUser ClanUser
	// Check first argument
	if isUser(command[1]) {
		addUser = ClanUser{
			UserName: m.Mentions[0].Username,
			UserID:   m.Mentions[0].ID,
			Nickname: getNickname(g, s, m.Mentions[0].ID),
			DateTime: time.Now(),
		}
	} else {
		message = fmt.Sprintf("Whoah, not so sure about those arguments. EventsBot is confused :confounded:")
		message = fmt.Sprintf("%s\r\nFor help with adding a user to the naughty list, type the following:\r\n```%shelp addnaughtylist```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	// Check that current user has permissions
	if !hasRole(g, s, m, "EventsBotAdmin") {
		message = fmt.Sprintf("Yo yo yo. Back up a second dude. You don't have permissions to add users to the naughty list.\r\nIf you're not careful then EventsBot might just add you to the naughty list :point_up:")
		message = fmt.Sprintf("%s\r\nFor help with adding a user to the naughty list, type the following:\r\n```%shelp addnaughtylist```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	c := mongoSession.DB(fmt.Sprintf("ClanEvents%s", g.ID)).C("NaughtyList")
	filter := bson.M{"userName": addUser.UserName}
	_, err := c.Upsert(filter, addUser)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(":scream::scream::scream:Something very weird happened when trying to add %s to the naughty list. Sorry but EventsBot has no answers for you :cry:", addUser.DisplayName()))
		return
	}

	message = fmt.Sprintf("%s has been added to the naughty list :angry:", addUser.DisplayName())
	s.ChannelMessageSend(m.ChannelID, message)
}

// RemoveNaughty is used to remove a user from the naughty list
func RemoveNaughty(g *discordgo.Guild, s *discordgo.Session, m *discordgo.MessageCreate, command []string) {
	message := ""

	// Test for correct number of arguments
	if len(command) != 2 {
		message = fmt.Sprintf("Whoah, not so sure about those arguments. EventsBot is confused :thinking:")
		message = fmt.Sprintf("%s\r\nFor help with removing a user from the naughty list, type the following:\r\n```%shelp removenaughtylist```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	var removeUser ClanUser
	// Check first argument
	if isUser(command[1]) {
		removeUser = ClanUser{
			UserName: m.Mentions[0].Username,
			UserID:   m.Mentions[0].ID,
			Nickname: getNickname(g, s, m.Mentions[0].ID),
			DateTime: time.Now(),
		}
	} else {
		message = fmt.Sprintf("Whoah, not so sure about those arguments. EventsBot is confused :confounded:")
		message = fmt.Sprintf("%s\r\nFor help with removing a user from the naughty list, type the following:\r\n```%shelp removenaughtylist```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	// Check that current user has permissions
	if !hasRole(g, s, m, "EventsBotAdmin") {
		message = fmt.Sprintf("Yo yo yo. Back up a second dude. You don't have permissions to remove users from the naughty list.\r\nIf you're not careful then EventsBot might just add you to the naughty list :point_up:")
		message = fmt.Sprintf("%s\r\nFor help with removing a user from the naughty list, type the following:\r\n```%shelp removenaughtylist```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	c := mongoSession.DB(fmt.Sprintf("ClanEvents%s", g.ID)).C("NaughtyList")
	filter := bson.M{"userName": removeUser.UserName}
	info, err := c.RemoveAll(filter)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(":scream::scream::scream:Something very weird happened when trying to remove %s from the naughty list. Sorry but EventsBot has no answers for you :cry:", removeUser.DisplayName()))
		return
	}

	if info.Removed > 0 {
		message = fmt.Sprintf("%s has been removed from the naughty list. Are we cool now? :kissing_heart:", removeUser.DisplayName())
	} else {
		message = fmt.Sprintf("What are you talking about? %s is not on the naughty list. :shrug:", removeUser.DisplayName())
	}
	s.ChannelMessageSend(m.ChannelID, message)
}

// AddServer is used to register a Discord server for ClanEvents to be able to run service functions for that server
func AddServer(g *discordgo.Guild, s *discordgo.Session, m *discordgo.MessageCreate, command []string) {
	message := ""

	// Test for correct number of arguments
	if len(command) != 1 {
		message = fmt.Sprintf("Whoah, not so sure about those arguments. EventsBot is confused :thinking:")
		message = fmt.Sprintf("%s\r\nFor help with adding a server to EventsBot, type the following:\r\n```%shelp addserver```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	// Check that current user has permissions
	if !hasRole(g, s, m, "EventsBotAdmin") {
		message = fmt.Sprintf("Yo yo yo. Back up a second dude. You don't have permissions to register servers.\r\nEventsBot will not stand for this :point_up:")
		message = fmt.Sprintf("%s\r\nFor help with adding a server to EventsBot, type the following:\r\n```%shelp addserver```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	c1 := mongoSession.DB(fmt.Sprintf("ClanEvents")).C("Guilds")
	var guild Guild
	guild.ID = g.ID
	guild.Name = g.Name
	filter := bson.M{"discordId": guild.ID}
	_, err := c1.Upsert(filter, guild)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, ":scream::scream::scream:Something very weird happened when trying to register this server. Sorry but EventsBot has no answers for you :cry:")
		return
	}

	c2 := mongoSession.DB(fmt.Sprintf("ClanEvents%s", g.ID)).C("Config")
	var config ClanConfig
	config.DefaultChannel = m.ChannelID
	filter = bson.M{}
	_, err = c2.Upsert(filter, config)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, ":scream::scream::scream:Something very weird happened when trying to register this server. Sorry but EventsBot has no answers for you :cry:")
		return
	}

	c3 := mongoSession.DB(fmt.Sprintf("ClanEvents%s", g.ID)).C("Events")
	index := mgo.Index{
		Key:        []string{"eventId"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c3.EnsureIndex(index)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, ":scream::scream::scream:Something very weird happened when trying to register this server. Sorry but EventsBot has no answers for you :cry:")
		return
	}

	message = fmt.Sprintf("%s has been registered with EventsBot", guild.Name)
	s.ChannelMessageSend(m.ChannelID, message)
}

// AddTimeZone is used to add capabilities for a time zone to ClanEvents
func AddTimeZone(g *discordgo.Guild, s *discordgo.Session, m *discordgo.MessageCreate, command []string) {
	message := ""

	// Test for correct number of arguments
	if len(command) != 3 {
		message = fmt.Sprintf("Whoah, not so sure about those arguments. EventsBot is confused :thinking:")
		message = fmt.Sprintf("%s\r\nFor help with adding a time zone, type the following:\r\n```%shelp addtimezone```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	// Check that current user has permissions
	if !hasRole(g, s, m, "EventsBotAdmin") {
		message = fmt.Sprintf("Yo yo yo. Back up a second dude. You don't have permissions to add time zones.\r\nYou don't look like you're from Gallifrey either :point_up:")
		message = fmt.Sprintf("%s\r\nFor help with adding a time zone, type the following:\r\n```%shelp addtimezone```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	// Check if abbreviation has already been used
	c := mongoSession.DB(fmt.Sprintf("ClanEvents%s", g.ID)).C("TimeZones")
	var newTZ TimeZone
	err := c.Find(bson.M{"abbrev": command[1]}).One(&newTZ)
	if err == nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("The abbreviation, %s, is already registered for %s", newTZ.Abbrev, newTZ.Location))
		return
	}

	newTZ.Abbrev = command[1]
	newTZ.Location = command[2]

	newLoc, err := time.LoadLocation(newTZ.Location)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("EventsBot is not so sure about that location. %s. Maybe you should double check that.", newTZ.Location))
		return
	}
	_, err = time.ParseInLocation("02/01/2006 15:04", "24/10/1975 12:00", newLoc)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("EventsBot is not so sure about that location. %s. Maybe you should double check that.", newTZ.Location))
		return
	}

	err = c.Insert(newTZ)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, ":scream::scream::scream:Something very weird happened when trying to add the time zone. Sorry but EventsBot has no answers for you :cry:")
		return
	}

	message = fmt.Sprintf("When you hear the signal it will be exactly %s, well in your newly registered timezone, %s, that is. Congrats... I guess.", time.Now().In(newLoc).Format("15:04"), newTZ.Abbrev)

	s.ChannelMessageSend(m.ChannelID, message)
}

// RemoveTimeZone is used to remove a time zone from ClanEvents
func RemoveTimeZone(g *discordgo.Guild, s *discordgo.Session, m *discordgo.MessageCreate, command []string) {
	message := ""

	// Test for correct number of arguments
	if len(command) != 2 {
		message = fmt.Sprintf("Whoah, not so sure about those arguments. EventsBot is confused :thinking:")
		message = fmt.Sprintf("%s\r\nFor help with removing a time zone, type the following:\r\n```%shelp removetimezone```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	// Check that current user has permissions
	if !hasRole(g, s, m, "EventsBotAdmin") {
		message = fmt.Sprintf("Yo yo yo. Back up a second dude. You don't have permissions to remove time zones.\r\nYou don't look like you're from Gallifrey either :point_up:")
		message = fmt.Sprintf("%s\r\nFor help with adding a time zone, type the following:\r\n```%shelp addtimezone```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	// Remove time zone from TimeZones collection
	ctz := mongoSession.DB(fmt.Sprintf("ClanEvents%s", g.ID)).C("TimeZones")
	filter := bson.M{"abbrev": command[1]}
	info, err := ctz.RemoveAll(filter)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(":scream::scream::scream:Something very weird happened when trying to remove %s from the time zones. Sorry but EventsBot has no answers for you :cry:", command[1]))
		return
	}

	// Remove all role time zones referencing this time zone from RoleTimeZones collection
	crtz := mongoSession.DB(fmt.Sprintf("ClanEvents%s", g.ID)).C("RoleTimeZones")
	_, err = crtz.RemoveAll(filter)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(":scream::scream::scream:Something very weird happened when trying to remove %s from the time zones. Sorry but EventsBot has no answers for you :cry:", command[1]))
		return
	}

	if info.Removed > 0 {
		message = fmt.Sprintf("%s has been removed from the list of time zones. Your world just got smaller.", command[1])
	} else {
		message = fmt.Sprintf("Are you trying to glitch the universe? %s is not in the list of time zones. :shrug:", command[1])
	}
	s.ChannelMessageSend(m.ChannelID, message)
}

// ListTimeZones is used to display a list of registered time zones
func ListTimeZones(g *discordgo.Guild, s *discordgo.Session, m *discordgo.MessageCreate, command []string) {
	message := ""

	// Test for correct number of arguments
	if len(command) != 1 {
		message = fmt.Sprintf("Whoah, not so sure about those arguments. EventsBot is confused :thinking:")
		message = fmt.Sprintf("%s\r\nFor help with listing time zones, type the following:\r\n```%shelp listtimezones```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	ctz := mongoSession.DB(fmt.Sprintf("ClanEvents%s", g.ID)).C("TimeZones")
	var timezones []TimeZone
	err := ctz.Find(bson.M{}).Sort("abbrev").All(&timezones)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, ":scream::scream::scream:Something very weird happened when trying to list the time zones. Sorry but EventsBot has no answers for you :cry:")
		return
	}

	crtz := mongoSession.DB(fmt.Sprintf("ClanEvents%s", g.ID)).C("RoleTimeZones")
	var roletzs []ServerRoleTimeZone
	err = crtz.Find(bson.M{}).Sort("serverRole").All(&roletzs)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, ":scream::scream::scream:Something very weird happened when trying to list the time zones. Sorry but EventsBot has no answers for you :cry:")
		return
	}

	message = fmt.Sprintf("%sList of registered time zones for %s:\r\n", message, g.Name)
	if len(timezones) == 0 {
		message = fmt.Sprintf("%sUhm. Nope. There ain't none. :shrug:", message)
	} else {
		for _, tz := range timezones {
			tzLoc, err := time.LoadLocation(tz.Location)
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, ":scream::scream::scream:Something very weird happened when trying to list the time zones. Sorry but EventsBot has no answers for you :cry:")
				return
			}
			message = fmt.Sprintf("%s**%s**: %s, current time %s\r\n", message, tz.Abbrev, tz.Location, time.Now().In(tzLoc).Format("15:04"))
		}
	}

	if len(roletzs) > 0 {
		message = fmt.Sprintf("%s\r\nServer roles with time zones:\r\n", message)
		for _, roletz := range roletzs {
			message = fmt.Sprintf("%s%s: %s\r\n", message, roletz.RoleName, roletz.Abbrev)
		}
	}

	s.ChannelMessageSend(m.ChannelID, message)
}

// RoleTimeZone is used to associate a time zone with a server role
func RoleTimeZone(g *discordgo.Guild, s *discordgo.Session, m *discordgo.MessageCreate, command []string) {
	message := ""

	// Test for correct number of arguments
	if len(command) != 3 {
		message = fmt.Sprintf("Whoah, not so sure about those arguments. EventsBot is confused :thinking:")
		message = fmt.Sprintf("%s\r\nFor help with linking a time zone to a server role, type the following:\r\n```%shelp roletimezone```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	// Check that current user has permissions
	if !hasRole(g, s, m, "EventsBotAdmin") {
		message = fmt.Sprintf("Yo yo yo. Back up a second dude. You don't have permissions to link server roles to time zones.\r\nIf you're not careful then EventsBot might just add you to the naughty list :point_up:")
		message = fmt.Sprintf("%s\r\nFor help with linking server roles to time zones, type the following:\r\n```%shelp roletimezone```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	// Check that the server role exists
	roleName := command[1]
	found := false
	for _, gRole := range g.Roles {
		if fmt.Sprintf("<@&%s>", gRole.ID) == roleName {
			found = true
		}
	}
	if !found {
		message = fmt.Sprintf("Say what? %s? EventsBot doesn't know any such server role.", roleName)
		message = fmt.Sprintf("%s\r\nFor help with linking server roles to time zones, type the following:\r\n```%shelp roletimezone```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	// Check that the time zone exists
	tz := command[2]
	ctz := mongoSession.DB(fmt.Sprintf("ClanEvents%s", g.ID)).C("TimeZones")
	count, err := ctz.Find(bson.M{"abbrev": tz}).Count()
	if err != nil || count == 0 {
		message = fmt.Sprintf("Say what? %s? EventsBot doesn't know any such time zone.", tz)
		message = fmt.Sprintf("%s\r\nFor help with linking server roles to time zones, type the following:\r\n```%shelp roletimezone```", message, config.CommandPrefix)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	// Link time zone to role
	var srtz ServerRoleTimeZone
	srtz.RoleName = roleName
	srtz.Abbrev = tz
	crtz := mongoSession.DB(fmt.Sprintf("ClanEvents%s", g.ID)).C("RoleTimeZones")
	filter := bson.M{"serverRole": roleName}
	_, err = crtz.Upsert(filter, srtz)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(":scream::scream::scream:Something very weird happened when trying to link the %s timezone to the %s server role. Sorry but EventsBot has no answers for you :cry:", srtz.Abbrev, srtz.RoleName))
		return
	}

	message = fmt.Sprintf("Booyaa! EventsBot linked the %s time zone to the %s server role", srtz.Abbrev, srtz.RoleName)
	s.ChannelMessageSend(m.ChannelID, message)
}

func isUser(arg string) bool {
	return strings.HasPrefix(arg, "<@")
}

func isDate(arg string) bool {
	_, err := time.Parse("02/01/2006", arg)
	return err == nil
}

func getGuild(s *discordgo.Session, m *discordgo.MessageCreate) *discordgo.Guild {
	// Attempt to get the channel from the state.
	// If there is an error, fall back to the restapi
	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		channel, err = s.Channel(m.ChannelID)
		if err != nil {
			return nil
		}
	}

	// Attempt to get the g from the state,
	// If there is an error, fall back to the restapi.
	g, err := s.State.Guild(channel.GuildID)
	if err != nil {
		g, err = s.Guild(channel.GuildID)
		if err != nil {
			return nil
		}
	}

	return g
}

func hasRole(g *discordgo.Guild, s *discordgo.Session, m *discordgo.MessageCreate, role string) bool {
	roleID := ""
	for _, gRole := range g.Roles {
		if gRole.Name == role {
			roleID = gRole.ID
			break
		}
	}
	found := false
	for _, role := range getRoles(g, m) {
		if role.ID == roleID {
			found = true
			break
		}
	}

	return found
}

func getRoles(g *discordgo.Guild, m *discordgo.MessageCreate) []*discordgo.Role {
	var retval []*discordgo.Role

	roles := make(map[string]*discordgo.Role)
	for _, gRole := range g.Roles {
		roles[gRole.ID] = gRole
	}

	for _, member := range g.Members {
		if member.User.Username == m.Author.Username {
			for _, memberRole := range member.Roles {
				retval = append(retval, roles[memberRole])
			}
			break
		}
	}

	return retval
}

func getNickname(g *discordgo.Guild, s *discordgo.Session, userID string) string {
	guildMember, err := s.GuildMember(g.ID, userID)
	if err != nil {
		return ""
	}
	return guildMember.Nick
}

func getLocation(g *discordgo.Guild, s *discordgo.Session, m *discordgo.MessageCreate) (*time.Location, string) {
	retloc := defaultLocation
	retabbr := ""

	// Start by getting all time zones and server role time zones
	ctz := mongoSession.DB(fmt.Sprintf("ClanEvents%s", g.ID)).C("TimeZones")
	crtz := mongoSession.DB(fmt.Sprintf("ClanEvents%s", g.ID)).C("RoleTimeZones")
	var tzs []TimeZone
	var roletzs []ServerRoleTimeZone
	tzLookup := make(map[string]TimeZone)
	err := ctz.Find(bson.M{}).Sort("abbrev").All(&tzs)
	if err != nil {
		return retloc, retabbr
	}
	err = crtz.Find(bson.M{}).Sort("serverRole").All(&roletzs)
	if err != nil {
		return retloc, retabbr
	}

	for _, tz := range tzs {
		tzLookup[tz.Abbrev] = tz
	}

	// Get all roles for specified user
	memberRoles := getRoles(g, m)

	// Find highest order role for member which has a time zone linked to it
	highest := 0

	for _, roletz := range roletzs {
		for _, memberRole := range memberRoles {
			if fmt.Sprintf("<@&%s>", memberRole.ID) == roletz.RoleName {
				if memberRole.Position > highest {
					highest = memberRole.Position
					retloc, _ = time.LoadLocation(tzLookup[roletz.Abbrev].Location)
					retabbr = roletz.Abbrev
				}
			}
		}
	}

	return retloc, retabbr
}
