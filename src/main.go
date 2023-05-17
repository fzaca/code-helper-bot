package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Xukay101/code-helper-bot/src/config"
	"github.com/Xukay101/code-helper-bot/src/database"
	"github.com/Xukay101/code-helper-bot/src/utils"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	// Load enviroment variables
	err := godotenv.Load()
	utils.FatalOnError("Error loading .env file", err)

	// Check if database starts - $ go run *src/*.go -init-db
	initDBFlag := flag.Bool("init-db", false, "Iniciar la base de datos")
	flag.Parse()
	if *initDBFlag {
		database.InitDb()
		os.Exit(0)
	}

	// Load config
	config := config.GetConfig()

	// Init bot
	bot := startBot(config.Bot.Token)
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
