package commands

import "github.com/bwmarrin/discordgo"

func HandlePing(s *discordgo.Session, m *discordgo.MessageCreate, prefix string) {
	// If the message is "ping" reply with "Pong!"
	if m.Content == prefix+"ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
}
