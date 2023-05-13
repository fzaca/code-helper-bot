package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Xukay101/code-helper-bot/src/utils"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	// Load enviroment variables
	err := godotenv.Load()
	utils.FatalOnError("Error loading .env file", err)

	// Load config
	config := loadConfig()

	// Init bot
	bot := startBot(config.botToken)
	bot.AddHandler(messageCreate)
	bot.Identify.Intents = discordgo.IntentsGuildMessages

	err = bot.Open()
	utils.FatalOnError("Error opening connection", err)

	// Signal for interrup
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Close Discord session.
	bot.Close()
}

func startBot(token string) *discordgo.Session {
	bot, err := discordgo.New("Bot " + token)
	utils.FatalOnError("Error creating Discord session", err)
	return bot
}
