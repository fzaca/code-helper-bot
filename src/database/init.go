package database

import (
	"log"

	"github.com/Xukay101/code-helper-bot/src/utils"
)

func InitDb() {
	conn, err := GetDb()
	utils.FatalOnError("Error trying to initialize database", err)
	defer conn.Close()

	for _, instruction := range instructions {
		_, err := conn.Exec(instruction)
		utils.FatalOnError("Error starting the database, wrong instruction", err)
	}

	log.Print("The database started correctly.")
}
