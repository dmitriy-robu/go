package mysql

import (
	"database/sql"
	"fmt"
	"go-rust-drop/config"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

var (
	onceDBMySQL     sync.Once
	mysqlConnection *sql.DB
)

func GetMySQLConnection() (*sql.DB, error) {
	configMySQl := config.LoadConfig().MySQL

	onceDBMySQL.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			configMySQl.User,
			configMySQl.Password,
			configMySQl.Host,
			configMySQl.Port,
			configMySQl.DBName,
		)

		var err error
		mysqlConnection, err = sql.Open("mysql", dsn)
		if err != nil {
			mysqlConnection = nil
			log.Fatalf("Error opening MySQL connection: %v", err)
		}
	})

	if mysqlConnection == nil {
		return nil, errors.New("Failed to initialize MySQL connection")
	}

	return mysqlConnection, nil
}
