package models

import (
	"context"
	"fmt"
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
	MongoDbPortEnvVar       = "MONGO_DB_PORT"
	MongoDbUserEnvVar       = "MONGO_DB_USER"
	MongoDbPasswordEnvVar   = "MONGO_DB_PASSWORD"
	MongoDbParamatersEnvVar = "MONGO_DB_PARAMETERS"
	MongoDbHostEnvVar       = "MONGO_DB_HOST"
	MongoCollection         = "Sessions"
	MongoDb                 = "SessionsDb"
	MongoDbDefaultPort      = "27017"
	MongoProtocol           = "mongodb://"
)

func getHost() string {
	// Get the port to use
	port := os.Getenv(MongoDbPortEnvVar)
	if port == "" {
		port = MongoDbDefaultPort
	}

	// provide a username and password if they're set
	mongoUser := os.Getenv(MongoDbUserEnvVar)
	mongoPassword := os.Getenv(MongoDbPasswordEnvVar)
	authString := ""
	if mongoUser != "" && mongoPassword != "" {
		authString = fmt.Sprintf("%s:%s@", mongoUser, mongoPassword)
	}

	parameters := os.Getenv(MongoDbParamatersEnvVar)
	if parameters != "" {
		parameters = "/?" + parameters
	}

	host := os.Getenv(MongoDbHostEnvVar)
	if host == "" {
		host = "localhost"
	}

	connectionString := fmt.Sprintf("mongodb://%s%s:%s%s", authString, host, port, parameters)
	log.Printf("Host: %s", connectionString)
	return connectionString
}

/// Initializes the database to interact with the session database
func InitializeDatabaseHandler() SessionDatabaseHandler {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(getHost()))
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
	var filter = bson.M{"_id": uniqueName}
	var session Session
	err := mongo.Collection.FindOne(context.Background(), filter).Decode(&session)
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
