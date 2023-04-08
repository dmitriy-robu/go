package mysql

import (
	"fmt"
	"go-rust-drop/config/db"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

var (
	onceGorm       sync.Once
	gormConnection *gorm.DB
)

func GetGormConnection() (*gorm.DB, error) {
	configMySQl := db.SetMysqlConfig()

	onceGorm.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			configMySQl.User,
			configMySQl.Password,
			configMySQl.Host,
			configMySQl.Port,
			configMySQl.DBName,
		)

		var err error
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
