package repositories

import (
	"go-rust-drop/internal/api/database/mongodb"
	"go-rust-drop/internal/api/database/mysql"
	"gorm.io/gorm"
	"log"
)

var MysqlDB = getMysqlDb()
var MongoDB = getMongoDb()

func getMysqlDb() *gorm.DB {
	db, err := mysql.GetGormConnection()
	if err != nil {
		log.Fatalln("Error getting MySQL connection")
	}

	return db
}

func getMongoDb() *mongodb.Client {
	db, err := mongodb.GetMongoDBConnection()
	if err != nil {
		log.Fatalln("Error getting MongoDB connection")
	}

	return db
}
