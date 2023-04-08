package migrations

import (
	mysql "go-rust-drop/internal/api/database/mysql"
	"log"
)

type Migrations struct {
}

func (m Migrations) MigrateAll() {
	mysqlConnection, err := mysql.GetGormConnection()
	if err != nil {
		log.Fatalln(err)
		return
	}

	CreateUsersTable(mysqlConnection)
	CreateUserBalancesTable(mysqlConnection)
}
