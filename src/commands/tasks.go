package commands

import (
	"github.com/bwmarrin/discordgo"
)

func HandleTasks(s *discordgo.Session, m *discordgo.MessageCreate, prefix string) {
	// If the command is "tasks" reply with tasks system
	if m.Content == prefix+"tasks" {
		s.ChannelMessageSend(m.ChannelID, "Tasks!")
	}
}
