package main

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/corentings/stouuf-bot/commands"
	"github.com/corentings/stouuf-bot/database"

	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
)

var (
	Token   string
	GuildID string
)

func loadVar() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	Token = os.Getenv("TOKEN")
	GuildID = os.Getenv("GUILD_ID")
}

func init() {
	loadVar()
}

func main() {
	// Try to connect to the database
	if err := database.Connect(); err != nil {
		log.Panic("Can't connect database:", err.Error())
	}
	log.Println("Connected to database successfully")

	defer func() {
		fmt.Println("Disconnect from database")
		err := database.Mg.Client.Disconnect(context.TODO())
		if err != nil {
			return
		}
	}()

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Println("error creating Discord session,", err)
		return
	}

	dg.Identify.Presence = discordgo.GatewayStatusUpdate{
		Game: discordgo.Activity{
			Name: "By CorentinGS",
			Type: discordgo.ActivityTypeCompeting,
		},
	}

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		log.Println("error opening connection,", err)
		return
	}

	commandHandlers := commands.GetCommandHandlers()

	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	appCommands := commands.GetCommands()

	_, err = dg.ApplicationCommandBulkOverwrite(dg.State.User.ID, GuildID, appCommands)
	if err != nil {
		log.Panicf("Error overwriting commands: %v", err)
	}

	defer func(dg *discordgo.Session) {
		err := dg.Close()
		if err != nil {
			return
		}
	}(dg)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	if true {
		log.Println("Removing commands...")
		_, err := dg.ApplicationCommandBulkOverwrite(dg.State.User.ID, GuildID, nil)
		if err != nil {
			log.Panicf("Cannot delete a command: %v", err)
		}
		log.Println("Removed commands")
	}
	log.Println("Gracefully shutting down.")
}
