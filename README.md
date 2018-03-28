# ClanEventsBot
[![Go Report Card](https://goreportcard.com/badge/github.com/deefstes/ClanEventsBot)](https://goreportcard.com/report/github.com/deefstes/ClanEventsBot) 
> Open Source Discord Bot written in Go. This is a simple bot that provides some calendar functionality to a Discord server enabling users to create and sign up to events.
  
| Option | Information |
|:--: | :--: |
| [Discord Developers](https://discordapp.com/developers/applications/me) | Register a bot account with Discord! |
| [Discord Go](https://github.com/bwmarrin/discordgo) | DiscordGo Library by: bwmarrin |
| [Discord Go (Go Docs)](https://godoc.org/github.com/bwmarrin/discordgo) | Godocs collection for DiscordGo |
| [mgo](https://labix.org/mgo) | Rich MongoDB driver for Go |
  
## Configuration
The `ClanEventsBot.yaml` file contains three entries for configuration:
`Token` -> Discord Bot Token
`CommandPrefix` -> The character that the bot expects commands to be prefixed with
`MongoDB` -> The address of the MongoDB server that serves as the datastore for the bot

## Database
The bot uses a MongoDB database and expects a database named `ClanEvents` with a collection named `Events` to be defined.

## Commands
```
listevents
details
newevent
cancelevent
signup
leave
impersonate
unimpersonate
```