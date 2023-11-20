package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStore struct {
	client   *mongo.Client
	database *mongo.Database
}

func NewMongoDatabase() (*MongoStore, error) {
	creds := options.Credential{
		Username: "root",
		Password: "test",
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(creds))
	if err != nil {
		return nil, err
	}

	return &MongoStore{
		client:   client,
		database: client.Database("accord-chat"),
	}, err
}

func (s *MongoStore) Close() {
	if err := s.client.Disconnect(context.Background()); err != nil {
		return
	}
}

func (s *MongoStore) GetClient() *mongo.Client {
	return s.client
}

func (s *MongoStore) GetDB() *mongo.Database {
	return s.database
}

func (s *MongoStore) InsertOne(ctx context.Context, collection string, data []byte) (string, error) {
	// Insert it into the database.
	result, err := s.database.Collection(collection).InsertOne(ctx, data)
	if err != nil {
		return primitive.NilObjectID.Hex(), err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (s *MongoStore) GetById(ctx context.Context, collection string, id string) ([]byte, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Get the data by the ID field.
	result, err := s.database.Collection(collection).FindOne(ctx, bson.M{"_id": objId}).Raw()
	if err != nil {
		return nil, err
	}

	return result, nil
}
