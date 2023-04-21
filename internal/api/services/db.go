package services

import (
	"go-rust-drop/internal/api/database/mysql"
	"gorm.io/gorm"
	"log"
)

var MysqlDB = getMysqlDb()

func getMysqlDb() *gorm.DB {
	db, err := mysql.GetGormConnection()
	if err != nil {
		log.Fatalln("Error getting MySQL connection")
	}

	return db
}
