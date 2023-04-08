package mongodb

import (
	"context"
	"fmt"
	"go-rust-drop/config/db"
	"net/url"
	"sync"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	onceDBMongoDB   sync.Once
	mongoConnection *mongo.Client
)

type Client struct {
	Client   *mongo.Client
	Database string
}

func GetMongoDBConnection() (*Client, error) {
	mongodbConfig := db.SetMongoDBConfig()

	onceDBMongoDB.Do(func() {
		uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authMechanism=%s&authSource=%s",
			mongodbConfig.User,
			url.QueryEscape(mongodbConfig.Password),
			mongodbConfig.Host,
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

	mongoDBClient := &Client{
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

func (m *Client) GetCollection(collectionName string) *mongo.Collection {
	return m.Client.Database(m.Database).Collection(collectionName)
}
