package migrations

import (
	"go-rust-drop/internal/api/models"
	"gorm.io/gorm"
	"log"
)

func CreateUserBalancesTable(db *gorm.DB) {
	var (
		err   error
		table models.UserBalance
	)

	if !db.Migrator().HasTable(table) {
		err = db.AutoMigrate(table)
		if err != nil {
			log.Fatalln(err)
			return
		}
	}
}
