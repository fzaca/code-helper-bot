package utils

import "log"

func FatalOnError(message string, err error) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}
