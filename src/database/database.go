package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func getDb() (*sql.DB, error) {
	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		GetConfig().Database.User,
		GetConfig().Database.Password,
		GetConfig().Database.Host,
		GetConfig().Database.Port,
		GetConfig().Database.Name,
	))
	return conn, err
}
