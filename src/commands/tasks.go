package commands

import (
	"fmt"
	"log"
	"math/rand"
	"strings"

	"github.com/Xukay101/code-helper-bot/src/database"
	"github.com/Xukay101/code-helper-bot/src/models"
	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
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

		// Get db
		conn := database.GetDb()

		// Get flag
		if len(args) > 1 {

			switch args[1] {

			case "--add":
				addTask(s, m, args, conn)
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

func addTask(s *discordgo.Session, m *discordgo.MessageCreate, args []string, conn *gorm.DB) {
	// addTask handles the "--add" subcommand of the "tasks" command
	if len(args) < 3 {
		commandHelp(s, m)
		return
	}

	// generate random identification that is not repeated
	generateRandomCode := func(serverId string) string {
		// get only the code of the tasks in the server
		var tasks []models.Task
		conn.Select("Code").Where("server_id = ?", serverId).Find(&tasks)

		codes := make(map[string]bool)
		for _, task := range tasks {
			fmt.Println(task.Code)
			codes[task.Code] = true
		}

		var code string
		for {
			code = fmt.Sprintf("%08d", rand.Intn(90000000)+10000000)
			if !codes[code] {
				break
			}
		}

		return code
	}

	// get fields
	description := strings.SplitN(m.Content, " ", 3)[2]
	createdBy := m.Author.ID
	serverId := m.GuildID
	code := generateRandomCode(serverId)

	// create new task
	task := models.Task{
		Code:        code,
		Description: description,
		CreatedBy:   createdBy,
		ServerId:    serverId,
	}

	result := conn.Create(&task)
	if result.Error != nil {
		log.Println(result.Error)
	}
}
