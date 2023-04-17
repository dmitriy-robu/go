package migrations

import (
	"go-rust-drop/internal/api/models"
	"log"
)

func CreateReferralTiersTable() {
	var (
		table models.ReferralTier
	)

	if !MySQL.Migrator().HasTable(table) {
		err = MySQL.AutoMigrate(table)
		if err != nil {
			log.Fatalln(err)
			return
		}
	}
}
