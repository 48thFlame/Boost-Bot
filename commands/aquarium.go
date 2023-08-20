package commands

import (
	"fmt"
	"strings"

	"github.com/48thFlame/Boost-bot/discord"
	dg "github.com/bwmarrin/discordgo"
)

var fish = [...]string{"🐠", "🐟", "🐡"}
var special = [...]string{"🐙", "🐬", "🐋", "🦈", "🦑", "🦐", "🪼", "🧜🏻‍♀️"}
var floor = [...]string{"🪸", "🌿", "🌱", "🌾", "🦀", "🦞", "🐚", "🏰"}
var top = [...]string{"🐳", "🌊", "⛵", "🏊", "🏄", "🚣", "🚤", "🚢", "🛟", "⛴️", ""}
var aquariumSpace = "\u3000\u200b"
var bubble = "🫧"

const (
	aquariumWidth  = 12
	aquariumHeight = 6
)

func genAquarium() string {
	var aquarium strings.Builder

	aquarium.WriteString(strings.Join(top[:], ""))
	aquarium.WriteRune('\n')
	aquarium.WriteString(strings.Join(fish[:], ""))
	aquarium.WriteString(bubble)
	aquarium.WriteRune('\n')
	aquarium.WriteString(strings.Join(special[:], ""))
	aquarium.WriteRune('\n')

	aquarium.WriteString(strings.Join(floor[:], ""))
	aquarium.WriteRune('\n')

	// for rowI := 0; rowI < aquariumHeight; rowI++ {
	// 	for colI := 0; colI < aquariumWidth; colI++ {
	// 		r := rand.Intn(20)
	// 		// r := 0
	// 		if r == 0 {
	// 			aquarium.WriteString(fish[rand.Intn(len(fish))])
	// 		} else {
	// 			aquarium.WriteString(aquariumSpace)
	// 		}
	// 	}
	// 	aquarium.WriteString("\n")
	// }

	return aquarium.String()
}

func Aquarium(s *dg.Session, i *dg.InteractionCreate) {
	// embed := discord.NewEmbed().
	// 	SetupEmbed().
	// 	SetAuthor("attachment://aquarium.png", "Aquarium", "")

	// embed.SetDescription(genAquarium())

	// aquariumR, err := os.Open("./assets/aquarium.png")
	// if err != nil {
	// 	discord.Error(fmt.Errorf("error opening info.png: %v", err), s, i.Interaction)
	// 	return
	// }

	// boostR, err := os.Open("./assets/boost.png")
	// if err != nil {
	// 	discord.Error(fmt.Errorf("error opening boost.png: %v", err), s, i.Interaction)
	// 	return
	// }

	err := discord.InteractionRespond(
		s,
		i.Interaction,
		dg.InteractionResponseChannelMessageWithSource,
		&dg.InteractionResponseData{
			Content: genAquarium(),
			// Embeds: []*dg.MessageEmbed{embed.MessageEmbed},
			// Files:  []*dg.File{{Name: "aquarium.png", Reader: aquariumR}, {Name: "boost.png", Reader: boostR}},
		},
	)
	if err != nil {
		discord.Error(fmt.Errorf("error responding to aquarium command interaction: %v", err), s, i.Interaction)
		return
	}
}
