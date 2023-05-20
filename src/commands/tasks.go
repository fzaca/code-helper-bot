package commands

import (
	"fmt"
	"log"
	"math/rand"
	"strings"

	"github.com/Xukay101/code-helper-bot/src/database"
	"github.com/bwmarrin/discordgo"
)

func HandleTasks(s *discordgo.Session, m *discordgo.MessageCreate, prefix string) {
	/*
		If the command is "tasks" reply with tasks system

		flags:
			--add [task] // Add global task
			--delete [id-task] // Delete global task
			--list (optional)[user] // return global list of tasks or list of user
			--edit [id-task] [task] // Edit global task
			--assign [id-task] (optional)[user] // Assign task to a user
			--unassign [id-task] // Unassign task
			--done [id-task] // Complete task
			--info [id-task] // Return with task information
	*/

	// Send message with `embed` use `ChannelMessageSendEmbed`

	args := strings.Split(m.Content, " ")

	if args[0] == prefix+"tasks" {

		if len(args) > 1 {

			switch args[1] {

			case "--add":
				addTask(s, m, args)
			default:
				commandHelp(s, m)

			}

		} else {
			commandHelp(s, m)
		}

	}
}

func commandHelp(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Return embed message with command help
	embed := &discordgo.MessageEmbed{
		Title: "Tasks Help",
		Description: `
			Pass
		`,
	}
	s.ChannelMessageSendEmbedReply(m.ChannelID, embed, m.Reference())
}

func addTask(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	// addTask handles the "--add" subcommand of the "tasks" command
	if len(args) < 3 {
		commandHelp(s, m)
		return
	}

	conn, err := database.GetDb()
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	generateRandomID := func() string {
		// generate random identification that is not repeated
		rows, err := conn.Query("SELECT id FROM tasks WHERE server_id = ?", m.GuildID)
		if err != nil {
			log.Println(err)
			return ""
		}
		defer rows.Close()

		existingIDs := make(map[int]bool)

		for rows.Next() {
			var id int
			err := rows.Scan(&id)
			if err != nil {
				log.Println(err)
				return ""
			}
			existingIDs[id] = true
		}

		var id int
		for {
			id = rand.Intn(90000000) + 10000000
			if !existingIDs[id] {
				break
			}
		}

		return fmt.Sprintf("%08d", id)
	}

	qry := "INSERT INTO tasks (id, description, created_by, server_id) VALUES (?, ?, ?, ?)"
	id := generateRandomID()
	if id == "" {
		return
	}
	description := strings.SplitN(m.Content, " ", 3)[2]
	createdBy := m.Author.ID
	serverId := m.GuildID

	_, err = conn.Exec(qry, id, description, createdBy, serverId)
	if err != nil {
		log.Println(err)
		return
	}
}

// func addTask(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
// }
