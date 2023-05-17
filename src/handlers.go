package main

import (
	"github.com/Xukay101/code-helper-bot/src/commands"
	"github.com/Xukay101/code-helper-bot/src/config"
	"github.com/bwmarrin/discordgo"
)

var prefix = config.GetConfig().Bot.Prefix

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Commands
	commands.HandlePing(s, m, prefix)
	commands.HandleDiagram(s, m, prefix)
	commands.HandleTasks(s, m, prefix)
}
