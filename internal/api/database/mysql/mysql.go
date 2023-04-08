package mysql

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-rust-drop/config/db"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sync"

	"github.com/pkg/errors"
)

var (
	onceGorm       sync.Once
	gormConnection *gorm.DB
)

func GetGormConnection() (*gorm.DB, error) {
	var err error

	if err = godotenv.Load(".env"); err != nil {
		log.Fatalln(err)
	}

	configMySQl := db.SetMysqlConfig()

	onceGorm.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			configMySQl.User,
			configMySQl.Password,
			configMySQl.Host,
			configMySQl.Port,
			configMySQl.DBName,
		)

		gormConnection, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			gormConnection = nil
			log.Fatalf("Error opening GORM connection: %v", err)
		}
	})

	if gormConnection == nil {
		return nil, errors.New("Failed to initialize GORM connection")
	}

	return gormConnection, nil
}
