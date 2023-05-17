package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Xukay101/code-helper-bot/src/config"
	"github.com/Xukay101/code-helper-bot/src/utils"
	_ "github.com/go-sql-driver/mysql"
)

func GetDb() (*sql.DB, error) {
	connString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=True",
		config.GetConfig().Database.User,
		config.GetConfig().Database.Password,
		config.GetConfig().Database.Host,
		config.GetConfig().Database.Port,
		config.GetConfig().Database.Name,
	)

	conn, err := sql.Open("mysql", connString)
	utils.FatalOnError("Error getting database", err)

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func initDb() {
	conn, err := GetDb()
	utils.FatalOnError("Error trying to initialize database", err)
	defer conn.Close()

	for _, instruction := range instructions {
		_, err := conn.Exec(instruction)
		utils.FatalOnError("Error starting the database, wrong instruction", err)
	}

	log.Print("The database started correctly.")
}
