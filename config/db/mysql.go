package db

import "os"

type MySQLConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func SetMysqlConfig() MySQLConfig {
	return MySQLConfig{
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		DBName:   os.Getenv("MYSQL_DBNAME"),
	}
}
