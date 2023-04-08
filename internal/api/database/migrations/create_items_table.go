package migrations

import (
	"go-rust-drop/internal/api/models"
	"log"
)

func CreateItemsTable() {
	var (
		table models.Item
	)

	if !MySQL.Migrator().HasTable(table) {
		err = MySQL.AutoMigrate(table)
		if err != nil {
			log.Fatalln(err)
			return
		}
	}
}
