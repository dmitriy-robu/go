package mongodb

import (
	"context"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/mongo/mongodriver"
	"github.com/pkg/errors"
	"go-rust-drop/config/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/url"
	"os"
	"sync"
	"time"
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
	var (
		err           error
		mongodbConfig db.MongoDBConfig
		uri           string
		client        *mongo.Client
		mongoDBClient *Client
	)

	mongodbConfig = db.SetMongoDBConfig()

	onceDBMongoDB.Do(func() {
		uri = fmt.Sprintf("mongodb://%s:%s@%s/%s?authMechanism=%s&authSource=%s&replicaSet=%s",
			mongodbConfig.User,
			url.QueryEscape(mongodbConfig.Password),
			mongodbConfig.ReplicaHost,
			mongodbConfig.DBName,
			mongodbConfig.AuthMechanism,
			mongodbConfig.AuthDatabase,
			mongodbConfig.ReplicaSet,
		)

		client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
		if err != nil {
			mongoConnection = nil
		} else {
			mongoConnection = client
		}
	})

	if mongoConnection == nil {
		return nil, errors.New("Failed to connect to MongoDB")
	}

	mongoDBClient = &Client{
		Client:   mongoConnection,
		Database: mongodbConfig.DBName,
	}

	return mongoDBClient, nil
}

func GetCollectionByName(collectionName string) (*mongo.Collection, error) {
	var (
		err           error
		mongoDBClient *Client
	)

	mongoDBClient, err = GetMongoDBConnection()
	if err != nil {
		return nil, errors.Wrap(err, "Error getting MongoDB connection")
	}

	return mongoDBClient.GetCollection(collectionName), nil
}

func (m *Client) GetCollection(collectionName string) *mongo.Collection {
	return m.Client.Database(m.Database).Collection(collectionName)
}

func InitMongoSessionStore() (sessions.Store, error) {
	var (
		secretKey  string
		collection *mongo.Collection
		store      sessions.Store
		err        error
	)

	secretKey = os.Getenv("SESSION_SECRET")
	if secretKey == "" {
		return nil, errors.New("key is not set")
	}

	collectionName := "sessions"

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection, err = GetCollectionByName(collectionName)
	if err != nil {
		return nil, err
	}

	store = mongodriver.NewStore(collection, 3600, true, []byte(secretKey))

	return store, nil
}
