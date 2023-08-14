package commands

import (
	"fmt"
	"os"

	"github.com/48thFlame/Boost-bot/discord"
	dg "github.com/bwmarrin/discordgo"
)

const (
	connect4APIUrl = "http://35.238.73.100:8080/connect4"
)

func Connect4(s *dg.Session, i *dg.InteractionCreate) {
	// 	// init vars
	subCommands := i.ApplicationCommandData().Options
	cmdName := subCommands[0].Name
	// 	id := discord.GetInteractionUser(i.Interaction).ID
	interaction := i.Interaction

	embed := discord.NewEmbed().
		SetupEmbed().
		SetAuthor("attachment://connect4.png", "Connect 4", "").
		SetDescription(cmdName)

	boostR, err := os.Open("./assets/boost.png")
	if err != nil {
		discord.Error(fmt.Errorf("error opening boost.png: %v", err), s, interaction)
	}

	connectR, err := os.Open("./assets/connect4.png")
	if err != nil {
		discord.Error(fmt.Errorf("error opening connect4.png: %v", err), s, interaction)
	}

	err = discord.InteractionRespond(
		s,
		interaction,
		discord.InstaMessage,
		&dg.InteractionResponseData{
			Embeds: []*dg.MessageEmbed{embed.MessageEmbed},
			Files:  []*dg.File{{Name: "boost.png", Reader: boostR}, {Name: "connect4.png", Reader: connectR}},
		},
	)
	if err != nil {
		discord.Error(fmt.Errorf("error responding to connect4 command interaction: %v", err), s, interaction)
	}
}
