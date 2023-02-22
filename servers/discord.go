package servers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"os"
)

var discordToken = "Token Goes Here" // Not actually. Your token should exist in the environment, or better yet, in some form of secrets. You should not store your token anywhere GitHub can access it! Include that location in your .gitignore!

// We want to use the following line in the powershell terminal:
// go build tutorial.go; ./tutorial.exe
// Don't forget to use .exe not .go for the executable, silly!

func main() {
	log.Trace("discord handler booting up")
	// TODO: Create Discord server-API object

	// Create Discord session
	session, err := discordgo.New("Bot " + discordToken)
	defer session.Close()
	if err != nil {
		log.WithError(err).Fatal("failed to create Discord session")
	}

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

func main() {
	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
