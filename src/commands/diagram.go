package commands

import "github.com/bwmarrin/discordgo"

func HandleDiagram(s *discordgo.Session, m *discordgo.MessageCreate, prefix string) {
	// If the message is "diagram" reply with diagram
	if m.Content == prefix+"diagram" {
		s.ChannelMessageSend(m.ChannelID, "diagrama!")
	}
}
