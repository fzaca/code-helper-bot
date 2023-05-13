package main

import "os"

type Config struct {
	botToken  string
	botPrefix string
}

func loadConfig() Config {
	return Config{
		botToken:  os.Getenv("DISCORD_TOKEN"),
		botPrefix: "code!",
	}
}
