package migrations

import (
	"go-rust-drop/internal/api/models"
	"log"
)

func CreateReferralTransactionsTable() {
	var (
		table models.ReferralTransaction
	)

	if !MySQL.Migrator().HasTable(table) {
		err = MySQL.AutoMigrate(table)
		if err != nil {
			log.Fatalln(err)
			return
		}
	}
}
