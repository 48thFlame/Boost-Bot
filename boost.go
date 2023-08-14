package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/48thFlame/Boost-bot/commands"
	"github.com/48thFlame/Boost-bot/discord"
)

func init() {
	// rand.Seed(time.Now().UnixNano())
	log.Default().SetOutput(os.Stdout)
}

func main() {
	err := run()
	if err != nil {
		log.Fatalf("Error running bot, with error:\n%v\n", err)
	}
}

func run() (err error) {
	config, err := loadConfig()
	if err != nil {
		return fmt.Errorf("error loading config: %v", err)
	}

	commands.FeedbackChannelId = config.FeedbackChannelId

	bot, err := discord.NewBot("./TOKEN.txt", config.PyInterpreterName, config.PyFilePath, commands.ExportCommands())
	if err != nil {
		return fmt.Errorf("error creating bot: %v", err)
	}

	err = bot.Start()
	if err != nil {
		return fmt.Errorf("error starting bot: %v", err)
	}
	defer bot.S.Close()

	// Wait for a quit signal to quit
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	log.Println("Gracefully shutting down...")

	return nil
}
