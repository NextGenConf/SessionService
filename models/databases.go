package models

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/// Interface to interact with a session database.
type SessionDatabaseHandler interface {
	GetAllSessions() ([]Session, error)
	GetSession(string) (Session, error)
	AddSession(Session) error
}

const (
	MongoDbHostEnvVar = "MONGO_DB_HOST"
	MongoCollection   = "Sessions"
	MongoDb           = "SessionsDb"
)

/// Initializes the database to interact with the session database
func InitializeDatabaseHandler() SessionDatabaseHandler {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv(MongoDbHostEnvVar)))
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
func (mongo *MongoDbHandler) GetAllSessions() ([]Session, error) {
	cursor, err := mongo.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Printf("Failed to read from db: %s", err.Error())
		return nil, err
	}

	var results []Session
	for cursor.Next(context.Background()) {
		var session Session
		err = cursor.Decode(&session)
		if err != nil {
			log.Printf("Failed to decode session: %s", err.Error())
			return nil, err
		}

		results = append(results, session)
	}

	return results, nil
}

/// Retrieve a session by the specific name
func (mongo *MongoDbHandler) GetSession(uniqueName string) (Session, error) {
	var filter = bson.M{"UniqueName": uniqueName}
	var session Session
	err := mongo.Collection.FindOne(context.Background(), filter).Decode(session)
	if err != nil {
		log.Printf("Failed to read session from db: %s", err.Error())
	}

	return session, nil
}

/// Adds a session to the database
func (mongo *MongoDbHandler) AddSession(session Session) error {
	data, err := bson.Marshal(session)
	if err != nil {
		log.Printf("Failed to serialize json: %s", err.Error())
		return err
	}

	_, err = mongo.Collection.InsertOne(context.Background(), data)
	if err != nil {
		log.Printf("Failed to insert new session: %s", err.Error())
		return err
	}

	return nil
}
