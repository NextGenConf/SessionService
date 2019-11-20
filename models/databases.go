package models

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/// Interface to interact with a session database.
type SessionDatabaseHandler interface {
	GetAllSessions() []Session
	GetSession(string) Session
}

const (
	MongoDbHost     = "mongodb://localhost:27018"
	MongoCollection = "Sessions"
	MongoDb         = "SessionsDb"
)

/// Initializes the database to interact with the session database
func InitializeDatabaseHandler() SessionDatabaseHandler {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoDbHost))
	if err != nil {
		panic("Failed to connect to Mongo DB")
	}

	handler := &MongoDbHandler{
		Collection: client.Database(MongoDb).Collection(MongoCollection),
	}
	return handler
}

/// A struct to manage a MongoDb session database
type MongoDbHandler struct {
	Collection *mongo.Collection
}

/// Retrieves all sessions from the database
func (mongo *MongoDbHandler) GetAllSessions() []Session {
	var sessions []Session
	cursor, err := mongo.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Printf("Failed to read from db: %s", err.Error())
		return make([]Session, 10)
	}

	cursor.All(context.Background(), sessions)
	return sessions
}

/// Retrieve a session by the specific name
func (mongo *MongoDbHandler) GetSession(uniqueName string) Session {
	var filter = bson.M{"UniqueName": uniqueName}
	var session Session
	mongo.Collection.FindOne(context.Background(), filter).Decode(session)
	return session
}
