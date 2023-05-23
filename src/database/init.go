package database

import (
	"log"
)

func InitDb() {
	conn := GetDb()

	for _, instruction := range instructions {
		conn.Exec(instruction)
	}

	log.Print("The database started correctly.")
}
