package commands

import (
	"github.com/bwmarrin/discordgo"
)

func HandlePing(s *discordgo.Session, m *discordgo.MessageCreate, prefix string) {
	// If the message is "ping" reply with "Pong!"
	if m.Content == prefix+"ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
}

// This is example

// func HandlePing(s *discordgo.Session, m *discordgo.MessageCreate, prefix string) {
// 	// If the message is "ping" reply with "Pong!"
// 	if m.Content == prefix+"ping" {
// 		conn, err := database.GetDb()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		defer conn.Close()
//
// 		row := conn.QueryRow("SELECT id, email, username, edad FROM Usuario WHERE id = ?", 2)
//
// 		var id int64
// 		var email string
// 		var username string
// 		var edad int64
//
// 		row.Scan(&id, &email, &username, &edad)
//
// 		message := fmt.Sprintf("Row:", id, email, username, edad)
// 		s.ChannelMessageSend(m.ChannelID, message)
//
// 		// s.ChannelMessageSend(m.ChannelID, "Pong!")
// 	}
// }
