package servers

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

// We want to use the following line in the powershell terminal:
// go build tutorial.go; ./tutorial.exe
// Don't forget to use .exe not .go for the executable, silly!

func main() {
	discordToken := "Token Goes Here" // Not actually. Your token should exist in the environment, or better yet, in some form of secrets. You should not store your token anywhere GitHub can access it! Include that location in your .gitignore!

	log.Trace("discord handler booting up")
	// TODO: Create Discord server-API object

	// Create Discord session
	session, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		log.WithError(err).Fatal("failed to create Discord session")
	}
	defer session.Close()

	// Actually open connection to Discord
	err = session.Open()
	if err != nil {
		log.WithError(err).Fatal("failed to open Discord connection")
	}

	// Wait until close
	// This is a good time to open the gRPC server(?)

	// TODO: Take message, parse into protobuf, send to backend // should be in a handler?
	// TODO: Send or act on response from backend // should be in a handler?
}
