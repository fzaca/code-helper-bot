package main

import "github.com/bwmarrin/discordgo"

var prefix = loadConfig().botPrefix

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// If the message is "ping" reply with "Pong!"
	if m.Content == prefix+"ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
}
