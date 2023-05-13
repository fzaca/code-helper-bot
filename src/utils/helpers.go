package utils

import "log"

func PrintError(message string, err error) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}
