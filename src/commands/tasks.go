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
			--info [id-task] // Return with task information
			--list (optional)[user] // return global list of tasks or list of user
			--delete [id-task] // Delete global task
			--edit [id-task] [task] // Edit global task
			--assign [id-task] (optional)[user] // Assign task to a user
			--unassign [id-task] // Unassign task
			--done [id-task] // Complete task
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
			case "--info":
				getInfo(s, m, args, conn)
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
	// AddTask handles the "--add" subcommand of the "tasks" command
	if len(args) < 3 {
		commandHelp(s, m)
		return
	}

	// Generate random identification that is not repeated
	generateRandomCode := func(serverId string) string {
		// Get only the code of the tasks in the server
		var tasks []models.Task
		conn.Select("Code").Where("server_id = ?", serverId).Find(&tasks)

		codes := make(map[string]bool)
		for _, task := range tasks {
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

	// Get fields
	description := strings.SplitN(m.Content, " ", 3)[2]
	createdBy := m.Author.ID
	serverId := m.GuildID
	code := generateRandomCode(serverId)

	// Create new task
	task := models.Task{
		Code:        code,
		Description: description,
		CreatedBy:   createdBy,
		ServerId:    serverId,
	}

	err := conn.Create(&task).Error
	if err != nil {
		log.Println(err)
	}
}

func getInfo(s *discordgo.Session, m *discordgo.MessageCreate, args []string, conn *gorm.DB) {
	// Check args
	if len(args) < 3 {
		commandHelp(s, m)
		return
	}

	// Get task
	var task models.Task

	err := conn.Where("server_id = ?", m.GuildID).First(&task, "code = ?", args[2]).Error
	if err != nil {
		return // Gorm generated log automatic
	}

	// Get users and format fields
	emptyUser := &discordgo.User{
		Username: " ",
	}
	assignedTo, err := s.User(task.AssignedTo)
	if err != nil {
		assignedTo = emptyUser
	}
	createdBy, err := s.User(task.CreatedBy)
	if err != nil {
		createdBy = emptyUser
	}
	updatedBy, err := s.User(task.UpdatedBy)
	if err != nil {
		updatedBy = emptyUser
	}
	completed := "No"
	if task.Completed {
		completed = "Yes"
	}

	// Send embed
	embed := &discordgo.MessageEmbed{
		Title: fmt.Sprintf("Task %s", task.Code),
		Description: fmt.Sprintf(`
			Description: %s 
			AssignedTo: %s
			Completed: %v
			CreatedBy: %s
			UpdatedBy: %s
			CreatedAt: %+v
			UpdatedAt: %+v
		`,
			task.Description,
			assignedTo.Username,
			completed,
			createdBy.Username,
			updatedBy.Username,
			task.CreatedAt,
			task.UpdatedAt,
		),
	}
	s.ChannelMessageSendEmbedReply(m.ChannelID, embed, m.Reference())
}
