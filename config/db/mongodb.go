package db

import "os"

type MongoDBConfig struct {
	Url           string
	User          string
	Password      string
	Host          string
	Port          string
	DBName        string
	AuthDatabase  string
	AuthMechanism string
	ReplicaSet    string
	ReplicaHost   string
}

func SetMongoDBConfig() MongoDBConfig {
	return MongoDBConfig{
		User:          os.Getenv("MONGODB_USER"),
		Password:      os.Getenv("MONGODB_PASSWORD"),
		Host:          os.Getenv("MONGODB_HOST"),
		Port:          os.Getenv("MONGODB_PORT"),
		DBName:        os.Getenv("MONGODB_DBNAME"),
		AuthDatabase:  os.Getenv("MONGODB_AUTH_DATABASE"),
		AuthMechanism: os.Getenv("MONGODB_AUTH_MECHANISM"),
		ReplicaSet:    os.Getenv("MONGODB_REPLICA_SET"),
		ReplicaHost:   os.Getenv("MONGODB_REPLICA_HOST"),
	}
}
