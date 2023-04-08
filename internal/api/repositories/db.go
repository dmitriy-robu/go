package repositories

import (
	"go-rust-drop/internal/api/database/mysql"
	"gorm.io/gorm"
	"log"
)

var MysqlDB = GetDb()

func GetDb() *gorm.DB {
	db, err := mysql.GetGormConnection()
	if err != nil {
		log.Fatalln("Error getting MySQL connection")
	}

	return db
}
