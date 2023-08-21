package commands

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/48thFlame/Boost-bot/discord"
	dg "github.com/bwmarrin/discordgo"
)

var fishArr = [...]string{"🐠", "🐟", "🐡"}
var fishSpecialArr = [...]string{"🐙", "🐬", "🐋", "🦈", "🦑", "🦐", "🪼", "🧜🏻‍♀️"}
var floorArr = [...]string{"🪸", "🌿", "🌱", "🌾"}
var floorSpecialArr = [...]string{"🦀", "🦞", "🐚", "🏰"}
var topArr = [...]string{"🐳", "🌊", "⛵", "🏊", "🏄", "🚣", "🚤", "🚢", "🛟", "⛴️"}

var aquariumSpace = "\u3000 \u200b "

// var aquariumSpace = "\u200b \u200b \u200b"
var bubble = "🫧"

const (
	topThingWeight = 0.15

	floorThingWeight   = 0.75
	floorSpecialWeight = 0.15

	levelThingWeight   = 0.18
	levelSpecialWeight = 0.1
)

// const fishArrLen = len(fishArr)
// const specialArrLen = len(specialArr)
// const floorArrLen = len(floorArr)
// const topArrayLen = len(topArr)

func randomWeight(weight float64) bool {
	return rand.Float64() < weight
}

func randomItem(opt []string) string {
	optLen := len(opt)
	i := rand.Intn(optLen)

	return opt[i]
}

func randomItemWeight(opt []string, weight float64) string {
	notEmpty := randomWeight(weight)
	if !notEmpty {
		return aquariumSpace
	} else {
		return randomItem(opt)
	}
}

func genTop(width int) string {
	if width < 1 {
		return ""
	}

	var topStr strings.Builder

	for i := 0; i < width; i++ {
		thing := randomItemWeight(topArr[:], topThingWeight)

		topStr.WriteString(thing)
	}

	topStr.WriteRune('\n')

	return topStr.String()
}

func genFloor(width int) string {
	if width < 1 {
		return ""
	}

	var floorStr strings.Builder

	for i := 0; i < width; i++ {
		var thing string

		shouldThing := randomWeight(floorThingWeight)
		if shouldThing {
			special := randomWeight(floorSpecialWeight)
			if special {
				thing = randomItem(floorSpecialArr[:])
			} else {
				thing = randomItem(floorArr[:])
			}
		} else {
			thing = aquariumSpace
		}

		floorStr.WriteString(thing)
	}

	floorStr.WriteRune('\n')

	return floorStr.String()
}

func genLevel(width int) string {
	if width < 1 {
		return ""
	}

	var levelStr strings.Builder

	for i := 0; i < width; i++ {
		var thing string

		shouldThing := randomWeight(levelThingWeight)
		if shouldThing {
			special := randomWeight(levelSpecialWeight)
			if special {
				thing = randomItem(fishSpecialArr[:])
			} else {
				thing = randomItem(fishArr[:])
			}
		} else {
			thing = aquariumSpace
		}

		levelStr.WriteString(thing)
	}

	levelStr.WriteRune('\n')

	return levelStr.String()
}

const width = 14
const height = 9

func genAquarium() string {
	var aquarium strings.Builder

	aquarium.WriteString(genTop(width))
	for i := 0; i < height; i++ {
		aquarium.WriteString(genLevel(width))
	}
	aquarium.WriteString(genFloor(width))

	// aquarium.WriteString(strings.Join(top[:], ""))
	// aquarium.WriteRune('\n')
	// aquarium.WriteString(strings.Join(fish[:], ""))
	// aquarium.WriteString(bubble)
	// aquarium.WriteRune('\n')
	// aquarium.WriteString(strings.Join(special[:], ""))
	// aquarium.WriteRune('\n')

	// aquarium.WriteString(strings.Join(floor[:], ""))
	// aquarium.WriteRune('\n')

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
