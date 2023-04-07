package config

import (
	"os"
)

type Config struct {
	MySQL   MySQLConfig
	MongoDB MongoDBConfig
}

type MySQLConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

type MongoDBConfig struct {
	Url           string
	User          string
	Password      string
	Host          string
	Port          string
	DBName        string
	AuthDatabase  string
	AuthMechanism string
}

func LoadConfig() *Config {
	return &Config{
		MySQL: MySQLConfig{
			User:     os.Getenv("MYSQL_USER"),
			Password: os.Getenv("MYSQL_PASSWORD"),
			Host:     os.Getenv("MYSQL_HOST"),
			Port:     os.Getenv("MYSQL_PORT"),
			DBName:   os.Getenv("MYSQL_DBNAME"),
		},
		MongoDB: MongoDBConfig{
			Url:           "localhost",
			User:          os.Getenv("MONGODB_USER"),
			Password:      os.Getenv("MONGODB_PASSWORD"),
			Host:          os.Getenv("MONGODB_HOST"),
			Port:          os.Getenv("MONGODB_PORT"),
			DBName:        os.Getenv("MONGODB_DBNAME"),
			AuthDatabase:  "admin",
			AuthMechanism: "SCRAM-SHA-256",
		},
	}
}
