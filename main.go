package main

import (
	"RUSLANGO/commands"
	"RUSLANGO/events"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

func main() {

	token := os.Getenv("TOKEN")

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(events.OnMessage)
	dg.AddHandler(events.OnServerJoin)
	dg.AddHandler(events.OnServerLeave)
	dg.AddHandler(events.OnBotReady)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C like 1000 times in a row to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	anekTimer(sc, dg)

	// Cleanly close down the Discord session.
	dg.Close()
}

// In order for program to be killed with the signal input, we need this function
func anekTimer(done <-chan os.Signal, dg *discordgo.Session) {

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			if commands.CheckTimeForAnecdote() {
				anecdote := commands.GetRandomAnecdote()
				dg.ChannelMessageSend(commands.General_Chat_ID, anecdote+"\n ДАННЫЙ АНЕКДОТ ПРОСПОНСИРОВАН ОЛЕГОМ ЕРМОЛАЕВЫМ")
			}
		}
	}
}
