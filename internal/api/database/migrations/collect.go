package migrations

import (
	mysql "go-rust-drop/internal/api/database/mysql"
	"gorm.io/gorm"
	"log"
)

type Migrations struct {
}

var (
	MySQL *gorm.DB
	err   error
)

func (m Migrations) MigrateAll() {
	MySQL, err = mysql.GetGormConnection()

	if err != nil {
		log.Fatalln(err)
		return
	}

	CreateUsersTable()
	CreateUserBalancesTable()
	CreateItemsTable()
	CreateReferralTiersTable()
	CreateReferralsTable()
	CreateReferralTransactionsTable()
	CreateBoxTable()
	CreateBoxItemsTable()
}
