package utils

import (
	"log"
)

func FatalOnError(message string, err error) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}

func PanicOnError(message string, err error) {
	if err != nil {
		log.Panicf("%s: %v", message, err)
	}
}

func PrintOnError(message string, err error) bool {
	if err != nil {
		log.Printf("%s: %v", message, err)
		return true
	} else {
		return false
	}
}
