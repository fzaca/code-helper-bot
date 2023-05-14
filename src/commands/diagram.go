package commands

import (
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Xukay101/code-helper-bot/src/utils"
	"github.com/bwmarrin/discordgo"
)

var rules string = "```" +
	`You can use the command as follows:
	- To generate a UML diagram directly in the message:
		code!diagram plantuml
		// Replace plantuml with your code

	- To load a UML code from a text file:
		code!diagram -txt
		// Attach a text file with the UML code` + "```"

var errorMsg string = "Error, could not generate the diagram"

func HandleDiagram(s *discordgo.Session, m *discordgo.MessageCreate, prefix string) {
	// Get args
	args := strings.Split(m.Content, " ")

	// Check if the message is a valid diagram command
	if len(args) > 1 && args[0] == prefix+"diagram" {
		if len(args) >= 2 {
			// Handle -txt command
			if args[1] == "-txt" {
				handleDiagramTxt(s, m)
				return
			}
		}

		// Handle inline diagram command
		handleInlineDiagram(s, m)
		return
	} else {
		s.ChannelMessageSend(m.ChannelID, rules)
	}
}

func handleDiagramTxt(s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(m.Attachments) > 0 {
		attachment := m.Attachments[0]
		ext := filepath.Ext(attachment.Filename)
		// Handle attached .txt file
		if ext == ".txt" {
			// Get attachment content
			resp, err := http.DefaultClient.Get(attachment.URL)
			if utils.PrintOnError("Error downloading attachment", err) {
				s.ChannelMessageSend(m.ChannelID, errorMsg)
				return
			}

			// Convert response
			respBytes, err := ioutil.ReadAll(resp.Body)
			if utils.PrintOnError("Error to get attachment content", err) {
				s.ChannelMessageSend(m.ChannelID, errorMsg)
				return
			}

			// Create temporarily file
			filePath := "src/assets/temp/" + attachment.ID + ".txt"
			file, err := os.Create(filePath)
			defer os.Remove(filePath)
			if utils.PrintOnError("Error creating temporarily file", err) {
				s.ChannelMessageSend(m.ChannelID, errorMsg)
				return
			}
			defer file.Close()

			// Write temporarily file
			file.Write(respBytes)

			// Execute python script
			cmd := exec.Command("python", "-m", "plantuml", filePath)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			cmd.Run() // err := cmd.Run()
			// utils.PanicOnError("Error running python plantuml", err)

			// Send image
			imagePath := filePath[:len(filePath)-3] + "png"
			image, err := os.Open(imagePath)
			defer os.Remove(imagePath)
			if utils.PrintOnError("Error to get diagram image", err) {
				s.ChannelMessageSend(m.ChannelID, errorMsg)
				return
			}
			defer image.Close()
			filename := filepath.Base(imagePath)
			s.ChannelFileSend(m.ChannelID, filename, image)
			return
		}
	}

	// No attached file error
	s.ChannelMessageSend(m.ChannelID, rules)
}

func handleInlineDiagram(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Handle inline diagram command

	code := m.Content[strings.Index(m.Content, " "):] // Get code
	codeBytes := []byte(code)

	// Create temporarily file
	filePath := "src/assets/temp/" + m.ID + ".txt"
	file, err := os.Create(filePath)
	defer os.Remove(filePath)
	if utils.PrintOnError("Error creating temporarily file", err) {
		s.ChannelMessageSend(m.ChannelID, errorMsg)
		return
	}
	defer file.Close()

	// Write temporarily file
	file.Write(codeBytes)

	// Execute python script
	cmd := exec.Command("python", "-m", "plantuml", filePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	// Send image
	imagePath := filePath[:len(filePath)-3] + "png"
	image, err := os.Open(imagePath)
	defer os.Remove(imagePath)
	if utils.PrintOnError("Error to get diagram image", err) {
		s.ChannelMessageSend(m.ChannelID, errorMsg)
		return
	}
	defer image.Close()
	filename := filepath.Base(imagePath)
	s.ChannelFileSend(m.ChannelID, filename, image)
	return
}
