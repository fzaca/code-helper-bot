package main

import (
	"github.com/Xukay101/code-helper-bot/src/commands"
	"github.com/bwmarrin/discordgo"
)

var prefix = GetConfig().Bot.Prefix

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Commands
	commands.HandlePing(s, m, prefix)
	commands.HandleDiagram(s, m, prefix)
}
