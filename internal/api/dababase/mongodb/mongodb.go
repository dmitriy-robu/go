package mongodb

import (
	"context"
	"fmt"
	"go-rust-drop/config"
	"sync"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	onceDBMongoDB   sync.Once
	mongoConnection *mongo.Client
)

type MongoDBClient struct {
	Client   *mongo.Client
	Database string
}

func GetMongoDBConnection() (*MongoDBClient, error) {
	mongodbConfig := config.LoadConfig().MongoDB

	onceDBMongoDB.Do(func() {
		uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authMechanism=%s&authSource=%s",
			mongodbConfig.User,
			mongodbConfig.Password,
			mongodbConfig.Url,
			mongodbConfig.Port,
			mongodbConfig.DBName,
			mongodbConfig.AuthMechanism,
			mongodbConfig.AuthDatabase,
		)

		var err error
		client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
		if err != nil {
			mongoConnection = nil
		} else {
			mongoConnection = client
		}
	})

	if mongoConnection == nil {
		return nil, errors.New("Failed to connect to MongoDB")
	}

	mongoDBClient := &MongoDBClient{
		Client:   mongoConnection,
		Database: mongodbConfig.DBName,
	}

	return mongoDBClient, nil
}

func GetCollectionByName(collectionName string) (*mongo.Collection, error) {
	mongoDBClient, err := GetMongoDBConnection()
	if err != nil {
		return nil, errors.Wrap(err, "Error getting MongoDB connection")
	}

	return mongoDBClient.GetCollection(collectionName), nil
}

func (m *MongoDBClient) GetCollection(collectionName string) *mongo.Collection {
	return m.Client.Database(m.Database).Collection(collectionName)
}