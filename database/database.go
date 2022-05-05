package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/deefstes/ClanEventsBot/logging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

type Database struct {
	client *mongo.Client
	dbName string
	ctx    context.Context
}

var ErrNoDocuments = errors.New("database: no documents")

func NewDatabase(connString string) (*Database, error) {
	cs, err := connstring.ParseAndValidate(connString)
	if err != nil {
		return nil, fmt.Errorf("parsing database connection string: %w", err)
	}
	// if cs.Database == "" {
	// 	return nil, fmt.Errorf("getting database name from connection string")
	// }
	client, err := mongo.NewClient(options.Client().ApplyURI(connString))
	if err != nil {
		return nil, fmt.Errorf("creating mongodb client: %w", err)
	}

	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("connecting to mongodb")
	}

	// Attempt connecting to the database
	log.Println(logging.LogEntry{
		Severity: "DEBUG",
		Message:  "attempting db connection",
	})
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Println(logging.LogEntry{
			Severity: "INFO",
			Message:  "unable to ping mongodb, continuing anyway",
		})
	} else {
		log.Println(logging.LogEntry{
			Severity: "DEBUG",
			Message:  "db connection successful",
		})
	}

	rsp := &Database{
		client: client,
		dbName: cs.Database,
		ctx:    ctx,
	}
	return rsp, nil
}

func (db *Database) Close() {
	db.client.Disconnect(db.ctx)
}

func (db *Database) Ping() (time.Duration, error) {
	t1 := time.Now()
	err := db.client.Ping(db.ctx, readpref.Primary())
	if err != nil {
		return time.Now().Sub(t1), fmt.Errorf("pinging mongodb: %w", err)
	}

	return time.Now().Sub(t1), nil
}

func (db *Database) AddGuild(guildID, guildName, defaultChannel string) (Guild, error) {
	c1 := db.client.Database(fmt.Sprintf("ClanEvents")).Collection("Guilds")
	var guild Guild
	guild.ID = guildID
	guild.Name = guildName
	filter := bson.M{"discordId": guild.ID}
	//guild.ObjectID = primitive.NilObjectID
	_, err := c1.ReplaceOne(
		db.ctx,
		filter,
		guild,
		options.Replace().SetUpsert(true),
	)
	if err != nil {
		return guild, fmt.Errorf("ClanEvents.Guilds.ReplaceOne(): %w", err)
	}

	c2 := db.client.Database(fmt.Sprintf("ClanEvents%s", guildID)).Collection("Config")
	var config ClanConfig
	config.DefaultChannel = defaultChannel
	filter = bson.M{}
	_, err = c2.ReplaceOne(
		db.ctx,
		filter,
		config,
		options.Replace().SetUpsert(true),
	)
	if err != nil {
		return guild, fmt.Errorf("ClanEvents%s.Config.ReplaceOne(): %w", guildID, err)
	}

	c3 := db.client.Database(fmt.Sprintf("ClanEvents%s", guildID)).Collection("Events")
	index := mongo.IndexModel{
		Keys: bson.M{
			"eventId": 1,
		},
		Options: options.Index().SetUnique(true).SetBackground(true).SetSparse(true),
	}

	_, err = c3.Indexes().CreateOne(db.ctx, index)
	if err != nil {
		return guild, fmt.Errorf("ClanEvents%s.Events.Indexes.CreateOne(): %w", guildID, err)
	}

	return guild, nil
}

func (db *Database) GetGuilds() ([]Guild, error) {
	var rsp []Guild
	collection := db.client.Database("ClanEvents").Collection("Guilds")
	cur, err := collection.Find(db.ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("ClanEvents.Guilds.Find(): %w", err)
	}
	if err = cur.All(db.ctx, &rsp); err != nil {
		return nil, fmt.Errorf("decoding guilds: %w", err)
	}

	return rsp, nil
}

func (db *Database) GetClanConfig(guildID string) (ClanConfig, error) {
	var rsp ClanConfig
	c := db.client.Database(fmt.Sprintf("ClanEvents%s", guildID)).Collection("Config")
	rslt := c.FindOne(db.ctx, bson.D{})
	if rslt.Err() == mongo.ErrNoDocuments {
		return rsp, ErrNoDocuments
	}
	if rslt.Err() != nil {
		return rsp, fmt.Errorf("ClanEvents%s.Config.FindOne(): %w", guildID, rslt.Err())
	}
	err := rslt.Decode(&rsp)
	if err != nil {
		return rsp, fmt.Errorf("decoding config for guid: %w", err)
	}

	return rsp, nil
}

func (db *Database) AddTimeZone(guildID string, tz TimeZone) error {
	c := db.client.Database(fmt.Sprintf("ClanEvents%s", guildID)).Collection("TimeZones")
	_, err := c.InsertOne(db.ctx, tz)
	if err != nil {
		return fmt.Errorf("ClanEvents%s.TimeZones.InsertOne(): %w", guildID, err)
	}

	return nil
}

func (db *Database) DeleteTimeZone(guildID, tzAbbr string) error {
	ctz := db.client.Database(fmt.Sprintf("ClanEvents%s", guildID)).Collection("TimeZones")
	filter := bson.M{"abbrev": tzAbbr}
	info, err := ctz.DeleteMany(db.ctx, filter)
	if err != nil {
		return fmt.Errorf("ClanEvents%s.TimeZones.DeleteMany: %w", guildID, err)
	}
	if info.DeletedCount == 0 {
		return ErrNoDocuments
	}

	return nil
}

func (db *Database) GetTimeZones(guildID string) ([]TimeZone, error) {
	var rsp []TimeZone
	c := db.client.Database(fmt.Sprintf("ClanEvents%s", guildID)).Collection("TimeZones")

	sortopts := options.Find().SetSort(bson.D{{Key: "abbrev", Value: 1}})
	cur, err := c.Find(db.ctx, bson.D{}, sortopts)
	if err != nil {
		return nil, fmt.Errorf("ClanEvents%s.TimeZones.Find(): %w", guildID, err)
	}
	if err = cur.All(db.ctx, &rsp); err != nil {
		return nil, fmt.Errorf("decoding timezones: %w", err)
	}

	return rsp, nil
}

func (db *Database) GetTimeZone(guildID, timezone string) (*TimeZone, error) {
	c := db.client.Database(fmt.Sprintf("ClanEvents%s", guildID)).Collection("TimeZones")
	var tz TimeZone
	rslt := c.FindOne(db.ctx, bson.M{"abbrev": timezone})
	if rslt.Err() != nil {
		if rslt.Err() == mongo.ErrNoDocuments {
			return nil, ErrNoDocuments
		}
		return nil, fmt.Errorf("ClanEvents%s.TimeZones.FindOne(): %w", guildID, rslt.Err())
	}
	err := rslt.Decode(&tz)
	if err != nil {
		return nil, fmt.Errorf("decoding timezone: %w", err)
	}

	return &tz, nil
}

func (db *Database) AddRoleTimeZone(guildID, roleName, tzAbbr string) error {
	var srtz ServerRoleTimeZone
	srtz.RoleName = roleName
	srtz.Abbrev = tzAbbr
	crtz := db.client.Database(fmt.Sprintf("ClanEvents%s", guildID)).Collection("RoleTimeZones")
	filter := bson.M{"serverRole": roleName}
	_, err := crtz.ReplaceOne(
		db.ctx,
		filter,
		srtz,
		options.Replace().SetUpsert(true),
	)
	if err != nil {
		return fmt.Errorf("ClanEvents%s.RoleTimeZones.Replace(): %w", guildID, err)
	}

	return nil
}

func (db *Database) GetRoleTimeZones(guildID string) ([]ServerRoleTimeZone, error) {
	var rsp []ServerRoleTimeZone
	c := db.client.Database(fmt.Sprintf("ClanEvents%s", guildID)).Collection("RoleTimeZones")

	sortopts := options.Find().SetSort(bson.D{{Key: "serverRole", Value: 1}})
	cur, err := c.Find(db.ctx, bson.D{}, sortopts)
	if err != nil {
		return nil, fmt.Errorf("ClanEvents%s.RoleTimeZones.Find(): %w", guildID, err)
	}
	if err = cur.All(db.ctx, &rsp); err != nil {
		return nil, fmt.Errorf("decoding role timezones: %w", err)
	}

	return rsp, nil
}

func (db *Database) DeleteRoleTimeZones(guildID, tzAbbr string) error {
	crtz := db.client.Database(fmt.Sprintf("ClanEvents%s", guildID)).Collection("RoleTimeZones")
	filter := bson.M{"abbrev": tzAbbr}
	info, err := crtz.DeleteMany(db.ctx, filter)
	if err != nil {
		return fmt.Errorf("ClanEvents%s.RoleTimeZones.DeleteMany(): %w", guildID, err)
	}
	if info.DeletedCount == 0 {
		return ErrNoDocuments
	}

	return nil
}

func (db *Database) GetEvents(guildID, user string, date time.Time) ([]ClanEvent, error) {
	c := db.client.Database(fmt.Sprintf("ClanEvents%s", guildID)).Collection("Events")

	filter := bson.M{}
	filter["dateTime"] = bson.M{
		"$gte": time.Now().Add(-1 * time.Hour),
	}
	if user != "all" {
		filter["participants.userName"] = user
	}
	if !date.IsZero() {
		filter["dateTime"] = bson.M{
			"$gte": date,
			"$lt":  date.AddDate(0, 0, 1),
		}
	}

	var results []ClanEvent
	sortopts := options.Find().SetSort(bson.D{{Key: "dateTime", Value: 1}})
	cur, err := c.Find(db.ctx, filter, sortopts)
	if err != nil {
		return nil, fmt.Errorf("ClanEvents%s.Events.Find(): %w", guildID, err)
	}
	if err = cur.All(db.ctx, &results); err != nil {
		return nil, fmt.Errorf("decoding events: %w", err)
	}

	return results, nil
}

func (db *Database) GetEvent(guildID, eventID string) (*ClanEvent, error) {
	c := db.client.Database(fmt.Sprintf("ClanEvents%s", guildID)).Collection("Events")

	var event ClanEvent
	rslt := c.FindOne(db.ctx, bson.M{"eventId": eventID})
	if rslt.Err() != nil {
		if rslt.Err() == mongo.ErrNoDocuments {
			return nil, ErrNoDocuments
		}
		return nil, fmt.Errorf("ClanEvents%s.Events.FindOne(): %w", guildID, rslt.Err())
	}
	err := rslt.Decode(&event)
	if err != nil {
		return nil, fmt.Errorf("decoding event: %w", err)
	}

	return &event, nil
}

func (db *Database) NewEvent(guildID string, event ClanEvent) error {
	c := db.client.Database(fmt.Sprintf("ClanEvents%s", guildID)).Collection("Events")
	_, err := c.InsertOne(db.ctx, event)
	if err != nil {
		return fmt.Errorf("ClanEvents%s.Events.InsertOne(): %w", guildID, err)
	}

	return nil
}

func (db *Database) DeleteEvent(guildID, eventID string) error {
	c := db.client.Database(fmt.Sprintf("ClanEvents%s", guildID)).Collection("Events")
	_, err := c.DeleteOne(db.ctx, bson.M{"eventId": eventID})
	if err != nil {
		return fmt.Errorf("ClanEvents%s.Events.DeleteOne(): %w", guildID, err)
	}

	return nil
}

func (db *Database) UpdateEvent(guildID string, event ClanEvent) error {
	//event.ObjectID = primitive.NilObjectID
	c := db.client.Database(fmt.Sprintf("ClanEvents%s", guildID)).Collection("Events")
	rslt, err := c.ReplaceOne(db.ctx, bson.M{"eventId": event.EventID}, event, options.Replace().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("ClanEvents%s.Events.ReplaceOne(): %w", guildID, err)
	}
	if rslt.ModifiedCount == 0 && rslt.UpsertedCount == 0 {
		return fmt.Errorf("ClanEvents%s.Events.ReplaceOne() made 0 replacements/upserts: %w", guildID, err)
	}

	return nil
}

func (db *Database) ArchiveEvents(guildID string) error {
	filter := bson.M{}
	filter["dateTime"] = bson.M{
		"$lte": time.Now().Add(-1 * time.Hour),
	}
	filter["archived"] = bson.M{
		"$ne": true,
	}

	c := db.client.Database(fmt.Sprintf("ClanEvents%s", guildID)).Collection("Events")

	var results []ClanEvent
	sortopts := options.Find().SetSort(bson.D{{Key: "dateTime", Value: 1}})
	cur, err := c.Find(db.ctx, filter, sortopts)
	if err != nil {
		return fmt.Errorf("ClanEvents%s.Events.Find(): %w", guildID, err)
	}
	if err = cur.All(db.ctx, &results); err != nil {
		return fmt.Errorf("decoding unarchived events: %w", err)
	}

	for _, event := range results {
		upsertfilter := bson.M{"eventId": event.EventID}
		event.Archived = true
		event.EventID = fmt.Sprintf("%s_%s", time.Now().Format("060102150405"), event.EventID)
		//event.ObjectID = primitive.NilObjectID
		_, err := c.ReplaceOne(
			db.ctx,
			upsertfilter,
			event,
			options.Replace().SetUpsert(true),
		)
		if err != nil {
			return fmt.Errorf("ClanEvents%s.Events.ReplaceOne(): %w", guildID, err)
		}
	}

	return nil
}

func (db *Database) SetNaughtyListInterval(guildID string, interval int64, randFact float64) error {
	c := db.client.Database(fmt.Sprintf("ClanEvents%s", guildID)).Collection("Config")

	_, err := c.UpdateOne(
		db.ctx,
		bson.M{},
		bson.D{{Key: "$set",
			Value: bson.D{
				{Key: "insultInterval", Value: interval},
				{Key: "insultRndFact", Value: randFact},
			},
		}},
	)
	if err != nil {
		return fmt.Errorf("ClanEvents%s.Config.UpdateOne(): %w", guildID, err)
	}

	return nil
}

func (db *Database) AddNaughtyList(guildID string, user ClanUser) error {
	c := db.client.Database(fmt.Sprintf("ClanEvents%s", guildID)).Collection("NaughtyList")
	filter := bson.M{"userName": user.UserName}
	_, err := c.ReplaceOne(
		db.ctx,
		filter,
		user,
		options.Replace().SetUpsert(true),
	)
	if err != nil {
		return fmt.Errorf("ClanEvents%s.NaughtyList.ReplaceOne(): %w", guildID, err)
	}

	return nil
}

func (db *Database) RemoveNaughtyList(guildID string, user ClanUser) error {
	c := db.client.Database(fmt.Sprintf("ClanEvents%s", guildID)).Collection("NaughtyList")
	filter := bson.M{"userName": user.UserName}
	info, err := c.DeleteMany(db.ctx, filter)
	if err != nil {
		return fmt.Errorf("ClanEvents%s.NaughtyList.DeleteMany(): %w", guildID, err)
	}

	if info.DeletedCount == 0 {
		return ErrNoDocuments
	}

	return nil
}

func (db *Database) GetNaughtyList(guildID string) ([]ClanUser, error) {
	c := db.client.Database(fmt.Sprintf("ClanEvents%s", guildID)).Collection("NaughtyList")
	sortopts := options.Find().SetSort(bson.D{{Key: "userName", Value: 1}})

	var naughtyList []ClanUser

	cur, err := c.Find(db.ctx, bson.D{}, sortopts)
	if err != nil {
		return nil, fmt.Errorf("ClanEvents%s.NaughtyList.Find(): %w", guildID, err)
	}
	if err = cur.All(db.ctx, &naughtyList); err != nil {
		return nil, fmt.Errorf("decoding naughty list: %w", err)
	}

	return naughtyList, nil
}
