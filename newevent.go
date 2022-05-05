package main

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/deefstes/ClanEventsBot/database"
	"github.com/deefstes/ClanEventsBot/logging"
)

type eventState int

const (
	stateNew      eventState = 0
	stateDate     eventState = 1
	stateTime     eventState = 2
	stateTimeZone eventState = 3
	stateDuration eventState = 4
	stateTeamSize eventState = 5
	stateDone     eventState = 6
)

type developingEvent struct {
	MessageID      string
	TriggerMessage *discordgo.MessageCreate
	State          eventState
	Event          database.ClanEvent
	Committed      bool
}

//gocyclo:ignore
// ShowDevelopingEvent is used to display the progress of an interactive new event
func ShowDevelopingEvent(s *discordgo.Session, m *discordgo.MessageCreate, channel string, newEvent developingEvent) {
	message := ""

	// Get channel
	c, err := s.Channel(channel)
	if err != nil {
		s.ChannelMessageSend(channel, "EventsBot had trouble obtaining the channel information :no_mouth:")
		return
	}

	// Get guild variables
	gv, ok := guildVarsMap[c.GuildID]
	if !ok {
		s.ChannelMessageSend(channel, "EventsBot had trouble obtaining the guild information :no_mouth:")
		return
	}

	// Get time zone
	tzInfo := ""
	eventLocation := defaultLocation

	if newEvent.Event.TimeZone != "" {
		tz, ok := gv.tzByAbbr[newEvent.Event.TimeZone]
		if !ok {
			s.ChannelMessageSend(channel, "EventsBot had trouble interpreting the time zone information of this event. Are we anywhere near a worm hole perhaps? :no_mouth:")
			return
		}
		tzInfo = tz.Abbrev
		eventLocation, _ = time.LoadLocation(tz.Location)
	} else {
		if len(gv.timezones) == 1 {
			tz := gv.timezones[0]
			tzInfo = tz.Abbrev
			newEvent.Event.TimeZone = tz.Abbrev
			eventLocation, _ = time.LoadLocation(tz.Location)
		}
	}

	// Construct message
	message = "NEW EVENT"
	message = fmt.Sprintf("%s\r\n**Creator:** %s", message, newEvent.Event.Creator.Mention())
	message = fmt.Sprintf("%s\r\n**Name:** %s", message, newEvent.Event.Name)
	if newEvent.State >= stateNew {
		message = fmt.Sprintf("%s\r\n**Date:** %s", message, newEvent.Event.DateTime.In(eventLocation).Format("Mon 2 Jan 2006"))
	}
	if newEvent.State >= stateTime {
		message = fmt.Sprintf("%s\r\n**Time:** %s", message, newEvent.Event.DateTime.In(eventLocation).Format("15:04"))
	}
	if newEvent.State > stateTimeZone { // note that this is >, not >= as we want to display the time zone only after it has been selected
		if newEvent.Event.TimeZone != "" {
			message = fmt.Sprintf("%s (%s)", message, tzInfo)
		}
	}
	if newEvent.State >= stateDuration {
		message = fmt.Sprintf("%s\r\n**Duration:** %d", message, newEvent.Event.Duration)
	}
	if newEvent.State >= stateTeamSize {
		message = fmt.Sprintf("%s\r\n**Team Size:** %d", message, newEvent.Event.TeamSize)
	}
	if newEvent.State >= stateDone {
		message = fmt.Sprintf("%s\r\n\r\nDoes the above appear correct?", message)
	}

	// Add appliccable reaction legend
	switch newEvent.State {
	case stateNew:
		fallthrough
	case stateDate:
		message = fmt.Sprintf("%s\r\n\r\n⏫ = Increase date by 1 month", message)
		message = fmt.Sprintf("%s\r\n🔼 = Increase date by 1 day", message)
		message = fmt.Sprintf("%s\r\n🔽 = Decrease date by 1 day", message)
		message = fmt.Sprintf("%s\r\n⏬ = Decrease date by 1 month", message)
		message = fmt.Sprintf("%s\r\n👍 = Continue", message)
		message = fmt.Sprintf("%s\r\n❌ = Cancel", message)
	case stateTime:
		message = fmt.Sprintf("%s\r\n\r\n⏪ = Decrease time by 1 hour", message)
		message = fmt.Sprintf("%s\r\n◀ = Decrease time by 10 minutes", message)
		message = fmt.Sprintf("%s\r\n▶ = Increase time by 10 minutes", message)
		message = fmt.Sprintf("%s\r\n⏩ = Increase time by 1 hour", message)
		message = fmt.Sprintf("%s\r\n👍 = Continue", message)
		message = fmt.Sprintf("%s\r\n❌ = Cancel", message)
	case stateTimeZone:
		message = fmt.Sprintf("%s\r\n\r\n Specify time zone", message)
		message = fmt.Sprintf("%s\r\n❌ = Cancel", message)
	case stateDuration:
		message = fmt.Sprintf("%s\r\n\r\n :one: - :nine: Specify duration (in hours)", message)
		message = fmt.Sprintf("%s\r\n❌ = Cancel", message)
	case stateTeamSize:
		message = fmt.Sprintf("%s\r\n\r\nSpecify team size:", message)
		if newEvent.Event.TeamSize < 10 {
			message = fmt.Sprintf("%s\r\n\r\n :one: - :nine: = 1 - 9", message)
			message = fmt.Sprintf("%s\r\n ▶ = More than 9", message)
		} else {
			message = fmt.Sprintf("%s\r\n\r\n :zero: - :nine: = %d0 - %d9", message, newEvent.Event.TeamSize/10, newEvent.Event.TeamSize/10)
			message = fmt.Sprintf("%s\r\n ◀ = Less than %d0", message, newEvent.Event.TeamSize/10)
			message = fmt.Sprintf("%s\r\n ▶ = More than %d9", message, newEvent.Event.TeamSize/10)
		}
		//message = fmt.Sprintf("%s\r\n👍 = Continue", message)
		message = fmt.Sprintf("%s\r\n❌ = Cancel", message)
	case stateDone:
		EditEvent(s, m, channel, newEvent.MessageID, "")
		return
		// message = fmt.Sprintf("%s\r\n✅ = OK", message)
		// message = fmt.Sprintf("%s\r\n❌ = Cancel", message)
		// message = fmt.Sprintf("%s\r\n🗓 = Back to Date", message)
		// message = fmt.Sprintf("%s\r\n🕑 = Back to Time", message)
		// message = fmt.Sprintf("%s\r\n🌍 = Back to Time Zone", message)
		// message = fmt.Sprintf("%s\r\n⏳ = Back to Duration", message)
		// message = fmt.Sprintf("%s\r\n👬 = Back to Team Size", message)
	default:
	}

	// Post or update message
	if newEvent.State == stateNew {
		newMsg, _ := s.ChannelMessageSend(channel, message)
		newEvent.MessageID = newMsg.ID
		gv.escrowEvents[newMsg.ID] = newEvent
	} else {
		s.ChannelMessageEdit(channel, newEvent.MessageID, "")
		s.ChannelMessageEdit(channel, newEvent.MessageID, message)
	}

	// Add appliccable reactions
	s.MessageReactionsRemoveAll(channel, newEvent.MessageID)
	switch newEvent.State {
	case stateNew:
		fallthrough
	case stateDate:
		s.MessageReactionAdd(channel, newEvent.MessageID, "⏫")
		s.MessageReactionAdd(channel, newEvent.MessageID, "🔼")
		s.MessageReactionAdd(channel, newEvent.MessageID, "🔽")
		s.MessageReactionAdd(channel, newEvent.MessageID, "⏬")
		s.MessageReactionAdd(channel, newEvent.MessageID, "👍")
		s.MessageReactionAdd(channel, newEvent.MessageID, "❌")
	case stateTime:
		s.MessageReactionAdd(channel, newEvent.MessageID, "⏪")
		s.MessageReactionAdd(channel, newEvent.MessageID, "◀")
		s.MessageReactionAdd(channel, newEvent.MessageID, "▶")
		s.MessageReactionAdd(channel, newEvent.MessageID, "⏩")
		s.MessageReactionAdd(channel, newEvent.MessageID, "👍")
		s.MessageReactionAdd(channel, newEvent.MessageID, "❌")
	case stateTimeZone:
		for emoji := range gv.tzByEmoji {
			s.MessageReactionAdd(channel, newEvent.MessageID, emoji)
		}
		s.MessageReactionAdd(channel, newEvent.MessageID, "❌")
	case stateDuration:
		s.MessageReactionAdd(channel, newEvent.MessageID, EmojiOne)
		s.MessageReactionAdd(channel, newEvent.MessageID, EmojiTwo)
		s.MessageReactionAdd(channel, newEvent.MessageID, EmojiThree)
		s.MessageReactionAdd(channel, newEvent.MessageID, EmojiFour)
		s.MessageReactionAdd(channel, newEvent.MessageID, EmojiFive)
		s.MessageReactionAdd(channel, newEvent.MessageID, EmojiSix)
		s.MessageReactionAdd(channel, newEvent.MessageID, EmojiSeven)
		s.MessageReactionAdd(channel, newEvent.MessageID, EmojiEight)
		s.MessageReactionAdd(channel, newEvent.MessageID, EmojiNine)
		s.MessageReactionAdd(channel, newEvent.MessageID, "❌")
	case stateTeamSize:
		if newEvent.Event.TeamSize > 9 {
			s.MessageReactionAdd(channel, newEvent.MessageID, "◀")
			s.MessageReactionAdd(channel, newEvent.MessageID, EmojiZero)
		}
		s.MessageReactionAdd(channel, newEvent.MessageID, EmojiOne)
		s.MessageReactionAdd(channel, newEvent.MessageID, EmojiTwo)
		s.MessageReactionAdd(channel, newEvent.MessageID, EmojiThree)
		s.MessageReactionAdd(channel, newEvent.MessageID, EmojiFour)
		s.MessageReactionAdd(channel, newEvent.MessageID, EmojiFive)
		s.MessageReactionAdd(channel, newEvent.MessageID, EmojiSix)
		s.MessageReactionAdd(channel, newEvent.MessageID, EmojiSeven)
		s.MessageReactionAdd(channel, newEvent.MessageID, EmojiEight)
		s.MessageReactionAdd(channel, newEvent.MessageID, EmojiNine)
		s.MessageReactionAdd(channel, newEvent.MessageID, "▶")
		//s.MessageReactionAdd(channel, newEvent.MessageID, "👍")
		s.MessageReactionAdd(channel, newEvent.MessageID, "❌")
	case stateDone:
		s.MessageReactionAdd(channel, newEvent.MessageID, "✅")
		s.MessageReactionAdd(channel, newEvent.MessageID, "❌")
		s.MessageReactionAdd(channel, newEvent.MessageID, "🗓")
		s.MessageReactionAdd(channel, newEvent.MessageID, "🕑")
		s.MessageReactionAdd(channel, newEvent.MessageID, "🌍")
		s.MessageReactionAdd(channel, newEvent.MessageID, "⏳")
		s.MessageReactionAdd(channel, newEvent.MessageID, "👬")
	default:
	}
}

//gocyclo:ignore
// ProcessReaction is used to respond to reactions added by the user to an interactive new event
func ProcessReaction(s *discordgo.Session, m *discordgo.MessageReactionAdd) {

	// Get channel
	c, err := s.Channel(m.MessageReaction.ChannelID)
	if err != nil {
		s.ChannelMessageSend(m.MessageReaction.ChannelID, "EventsBot had trouble obtaining the channel information :no_mouth:")
		return
	}

	// Find message in EscrowEvents
	gv, ok := guildVarsMap[c.GuildID]
	if !ok {
		return
	}
	event, ok := gv.escrowEvents[m.MessageID]
	if !ok {
		return
	}

	log.Println(logging.LogEntry{
		Severity: "DEBUG",
		Message:  fmt.Sprintf("%s reaction received for message %s", m.MessageReaction.Emoji.Name, event.MessageID),
	})

	// Respond to reaction based on state of developing event
	switch event.State {
	case stateNew:
		fallthrough
	case stateDate:
		event.State = stateDate
		switch m.MessageReaction.Emoji.Name {
		case "⏫":
			event.Event.DateTime = event.Event.DateTime.AddDate(0, 1, 0)
		case "🔼":
			event.Event.DateTime = event.Event.DateTime.AddDate(0, 0, 1)
		case "🔽":
			event.Event.DateTime = event.Event.DateTime.AddDate(0, 0, -1)
		case "⏬":
			event.Event.DateTime = event.Event.DateTime.AddDate(0, -1, 0)
		case "👍":
			if event.Committed {
				event.State = stateDone
			} else {
				event.State = stateTime
			}
		case "❌":
			delete(gv.escrowEvents, m.MessageID)
			s.ChannelMessageDelete(m.MessageReaction.ChannelID, m.MessageID)
			return
		}
	case stateTime:
		switch m.MessageReaction.Emoji.Name {
		case "⏪":
			event.Event.DateTime = event.Event.DateTime.Add(-1 * time.Hour)
		case "◀":
			event.Event.DateTime = event.Event.DateTime.Add(-10 * time.Minute)
		case "▶":
			event.Event.DateTime = event.Event.DateTime.Add(10 * time.Minute)
		case "⏩":
			event.Event.DateTime = event.Event.DateTime.Add(1 * time.Hour)
		case "👍":
			if event.Committed {
				event.State = stateDone
			} else {
				if event.Event.TimeZone != "" {
					event.State = stateDuration
				} else {
					event.State = stateTimeZone
				}
			}
		case "❌":
			delete(gv.escrowEvents, m.MessageID)
			s.ChannelMessageDelete(m.MessageReaction.ChannelID, m.MessageID)
			return
		}
	case stateTimeZone:
		timezone, ok := gv.tzByEmoji[m.MessageReaction.Emoji.Name]
		if ok {
			dtyr := event.Event.DateTime.Year()
			dtmo := event.Event.DateTime.Month()
			dtda := event.Event.DateTime.Day()
			dtho := event.Event.DateTime.Hour()
			dtmi := event.Event.DateTime.Minute()
			event.Event.TimeZone = timezone.Abbrev
			location, _ := time.LoadLocation(timezone.Location)
			event.Event.DateTime = time.Date(dtyr, dtmo, dtda, dtho, dtmi, 0, 0, location)
			if event.Committed {
				event.State = stateDone
			} else {
				event.State = stateDuration
			}
		}
		if m.MessageReaction.Emoji.Name == "❌" {
			delete(gv.escrowEvents, m.MessageID)
			s.ChannelMessageDelete(m.MessageReaction.ChannelID, m.MessageID)
			return
		}
	case stateDuration:
		if event.Committed {
			event.State = stateDone
		} else {
			event.State = stateTeamSize
		}
		switch m.MessageReaction.Emoji.Name {
		case EmojiOne:
			event.Event.Duration = 1
		case EmojiTwo:
			event.Event.Duration = 2
		case EmojiThree:
			event.Event.Duration = 3
		case EmojiFour:
			event.Event.Duration = 4
		case EmojiFive:
			event.Event.Duration = 5
		case EmojiSix:
			event.Event.Duration = 6
		case EmojiSeven:
			event.Event.Duration = 7
		case EmojiEight:
			event.Event.Duration = 8
		case EmojiNine:
			event.Event.Duration = 9
		case "❌":
			delete(gv.escrowEvents, m.MessageID)
			s.ChannelMessageDelete(m.MessageReaction.ChannelID, m.MessageID)
			return
		default:
			event.State = stateDuration
		}
	case stateTeamSize:
		event.State = stateDone
		baseSize := (event.Event.TeamSize / 10) * 10
		switch m.MessageReaction.Emoji.Name {
		case EmojiZero:
			event.Event.TeamSize = baseSize + 1
		case EmojiOne:
			event.Event.TeamSize = baseSize + 1
		case EmojiTwo:
			event.Event.TeamSize = baseSize + 2
		case EmojiThree:
			event.Event.TeamSize = baseSize + 3
		case EmojiFour:
			event.Event.TeamSize = baseSize + 4
		case EmojiFive:
			event.Event.TeamSize = baseSize + 5
		case EmojiSix:
			event.Event.TeamSize = baseSize + 6
		case EmojiSeven:
			event.Event.TeamSize = baseSize + 7
		case EmojiEight:
			event.Event.TeamSize = baseSize + 8
		case EmojiNine:
			event.Event.TeamSize = baseSize + 9
		case "◀":
			if baseSize >= 10 {
				event.Event.TeamSize = baseSize - 10
				if event.Event.TeamSize <= 0 {
					event.Event.TeamSize = 1
				}
			}
			event.State = stateTeamSize
		case "▶":
			event.Event.TeamSize = baseSize + 10
			event.State = stateTeamSize
		//case "👍":
		//	event.State = stateDone
		case "❌":
			delete(gv.escrowEvents, m.MessageID)
			s.ChannelMessageDelete(m.MessageReaction.ChannelID, m.MessageID)
			return
		default:
			event.State = stateTeamSize
		}
	case stateDone:
		switch m.MessageReaction.Emoji.Name {
		case "🗓":
			event.State = stateDate
		case "🕑":
			event.State = stateTime
		case "🌍":
			event.State = stateTimeZone
		case "⏳":
			event.State = stateDuration
		case "👬":
			event.State = stateTeamSize
		case "✅":
			CommitEvent(s, m.MessageReaction.ChannelID, event)
			fallthrough
		case "❌":
			delete(gv.escrowEvents, m.MessageID)
			s.ChannelMessageDelete(m.MessageReaction.ChannelID, m.MessageID)
			return
		}
	}

	gv.escrowEvents[m.MessageID] = event

	ShowDevelopingEvent(s, nil, m.MessageReaction.ChannelID, event)
}

// CommitEvent is used to move an event from Escrow to the DB
func CommitEvent(s *discordgo.Session, channelID string, newEvent developingEvent) {
	// Get channel
	channel, err := s.Channel(channelID)
	if err != nil {
		s.ChannelMessageSend(channelID, "EventsBot had trouble obtaining the channel information :no_mouth:")
		return
	}

	// Get guild
	guild, err := s.Guild(channel.GuildID)
	if err != nil {
		s.ChannelMessageSend(channelID, "EventsBot had trouble obtaining the guild information :no_mouth:")
		return
	}

	err = db.UpdateEvent(channel.GuildID, newEvent.Event)
	if err != nil {
		s.ChannelMessageSend(channelID, ":scream::scream::scream:Something very weird happened when trying to create this event. Sorry but EventsBot has no answers for you :cry:")
		return
	}

	if !newEvent.Committed {
		message := fmt.Sprintf("Woohoo! A new event has been created by %s. EventsBot is most pleased :ok_hand:", newEvent.Event.Creator.Mention())
		message = fmt.Sprintf("%s\r\nEvent ID: **%s**", message, newEvent.Event.EventID)
		message = fmt.Sprintf("%s\r\n\r\nTo sign up for this event, type the following:", message)
		message = fmt.Sprintf("%s\r\n```%ssignup %s```", message, config.CommandPrefix, newEvent.Event.EventID)
		s.ChannelMessageSend(channelID, message)

		signupCmd := []string{"signup", newEvent.Event.EventID}
		Signup(guild, s, newEvent.TriggerMessage, signupCmd)
	} else {
		message := fmt.Sprintf("Yeah boi! %s has successfully modified event %s. EventsBot is impressed :ok_hand:", newEvent.Event.Creator.Mention(), newEvent.Event.EventID)
		s.ChannelMessageSend(channelID, message)

		detailsCmd := []string{"details", newEvent.Event.EventID}
		Details(guild, s, newEvent.TriggerMessage, detailsCmd)
	}
}

// EditEvent is used to change the details of an event
func EditEvent(s *discordgo.Session, m *discordgo.MessageCreate, channelID string, messageID string, eventID string) {

	var devEvt developingEvent

	// Get channel
	c, err := s.Channel(channelID)
	if err != nil {
		s.ChannelMessageSend(channelID, "EventsBot had trouble obtaining the channel information :no_mouth:")
		return
	}

	// Find message in EscrowEvents
	gv, ok := guildVarsMap[c.GuildID]
	if !ok {
		return
	}
	_, ok = gv.escrowEvents[messageID]
	if !ok {
		// If no event is found in escrow for the specified message, it could mean that it's referring to an event already in the db and needs to be pulled from there
		event, err := db.GetEvent(c.GuildID, eventID)
		if err == errNoRecords {
			s.ChannelMessageSend(channelID, fmt.Sprintf("EventsBot could find no such event. Are you sure you got that Event ID of %s right? Them's finicky numbers. :grimacing:", eventID))
			return
		} else if err != nil {
			log.Println(logging.LogEntry{
				Severity: "ERROR",
				Message:  fmt.Sprintf("database: %+v", err),
			})
			s.ChannelMessageSend(channelID, ":scream::scream::scream:Something very weird happened when trying to edit this event. Sorry but EventsBot has no answers for you :cry:")
			return
		}

		newEvent := developingEvent{
			TriggerMessage: m,
			MessageID:      messageID,
			State:          stateDone,
			Event:          *event,
			Committed:      true,
		}
		gv.escrowEvents[messageID] = newEvent
		devEvt = newEvent
	}
	devEvt, ok = gv.escrowEvents[messageID]
	if !ok {
		s.ChannelMessageSend(channelID, "EventsBot had trouble interpreting the developing event. This is one of those things that should happen but then they do. :face_with_spiral_eyes:")
		return
	}

	// Get time zone
	tzInfo := ""
	eventLocation := defaultLocation

	if devEvt.Event.TimeZone != "" {
		tz, ok := gv.tzByAbbr[devEvt.Event.TimeZone]
		if !ok {
			s.ChannelMessageSend(channelID, "EventsBot had trouble interpreting the time zone information of this event. Are we anywhere near a worm hole perhaps? :no_mouth:")
			return
		}
		tzInfo = tz.Abbrev
		eventLocation, _ = time.LoadLocation(tz.Location)
	}

	// Construct message
	message := "EDIT EVENT"
	message = fmt.Sprintf("%s\r\n**Creator:** %s", message, devEvt.Event.Creator.Mention())
	message = fmt.Sprintf("%s\r\n**Name:** %s", message, devEvt.Event.Name)
	message = fmt.Sprintf("%s\r\n**Date:** %s", message, devEvt.Event.DateTime.In(eventLocation).Format("Mon 2 Jan 2006"))
	message = fmt.Sprintf("%s\r\n**Time:** %s", message, devEvt.Event.DateTime.In(eventLocation).Format("15:04"))
	if devEvt.Event.TimeZone != "" {
		message = fmt.Sprintf("%s (%s)", message, tzInfo)
	}
	message = fmt.Sprintf("%s\r\n**Duration:** %d", message, devEvt.Event.Duration)
	message = fmt.Sprintf("%s\r\n**Team Size:** %d", message, devEvt.Event.TeamSize)
	message = fmt.Sprintf("%s\r\n\r\nDoes the above appear correct?", message)
	message = fmt.Sprintf("%s\r\n✅ = OK", message)
	message = fmt.Sprintf("%s\r\n❌ = Cancel", message)
	message = fmt.Sprintf("%s\r\n🗓 = Change Date", message)
	message = fmt.Sprintf("%s\r\n🕑 = Change Time", message)
	message = fmt.Sprintf("%s\r\n🌍 = Change Time Zone", message)
	message = fmt.Sprintf("%s\r\n⏳ = Change Duration", message)
	message = fmt.Sprintf("%s\r\n👬 = Change Team Size", message)

	// Post or update message
	if messageID == "" {
		newMsg, _ := s.ChannelMessageSend(channelID, message)
		gv.escrowEvents[newMsg.ID] = devEvt
	} else {
		s.ChannelMessageEdit(channelID, messageID, "")
		s.ChannelMessageEdit(channelID, messageID, message)
	}

	// Add appliccable reactions
	s.MessageReactionsRemoveAll(channelID, devEvt.MessageID)
	s.MessageReactionAdd(channelID, devEvt.MessageID, "✅")
	s.MessageReactionAdd(channelID, devEvt.MessageID, "❌")
	s.MessageReactionAdd(channelID, devEvt.MessageID, "🗓")
	s.MessageReactionAdd(channelID, devEvt.MessageID, "🕑")
	s.MessageReactionAdd(channelID, devEvt.MessageID, "🌍")
	s.MessageReactionAdd(channelID, devEvt.MessageID, "⏳")
	s.MessageReactionAdd(channelID, devEvt.MessageID, "👬")
}
