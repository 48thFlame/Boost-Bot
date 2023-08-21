package commands

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/48thFlame/Boost-bot/discord"
	dg "github.com/bwmarrin/discordgo"
)

var fishArr = [...]string{"🐠 ", "🐟 ", "🐡 "}
var fishSpecialArr = [...]string{"🐙 ", "🐬 ", "🐋 ", "🦈 ", "🦑 ", "🦐 ", "🪼 ", "🧜🏻‍♀️ "}
var floorArr = [...]string{"🪸 ", "🌿 ", "🌱 ", "🌾 "}
var floorSpecialArr = [...]string{"🦀 ", "🦞 ", "🐚 ", "🏰 "}
var topArr = [...]string{"🐳 ", "🌊 ", "⛵ ", "🏊 ", "🏄 ", "🚣 ", "🚤 ", "🚢 ", "🛟 ", "⛴️ "}

var aquariumSpace = "\u3000\u3000"

var bubble = "🫧"

const (
	topThingWeight = 0.3

	floorThingWeight   = 0.8
	floorSpecialWeight = 0.18

	levelThingWeight   = 0.19
	levelSpecialWeight = 0.252

	bubbleWeight = 0.078

	defaultAquariumWidth  int64 = 10
	defaultAquariumHeight int64 = 7
)

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

		shouldBubble := randomWeight(bubbleWeight)

		if shouldBubble {
			thing = bubble
		} else {
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
		}

		levelStr.WriteString(thing)
	}

	levelStr.WriteRune('\n')

	return levelStr.String()
}

func genAquarium(width, height int) string {
	var aquarium strings.Builder

	aquarium.WriteString(genTop(width))
	for i := 0; i < height; i++ {
		aquarium.WriteString(genLevel(width))
	}
	aquarium.WriteString(genFloor(width))

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
	options := i.ApplicationCommandData().Options
	optionsLen := len(options)
	width := defaultAquariumWidth
	height := defaultAquariumHeight

	if optionsLen >= 2 {
		width = options[0].IntValue()
		height = options[1].IntValue()
	} else if optionsLen == 1 {
		firstOpt := options[0]
		switch firstOpt.Name {
		case "width":
			width = firstOpt.IntValue()
		case "height":
			height = firstOpt.IntValue()
		default:
			discord.Error(
				fmt.Errorf(
					"unrecognized options in aquarium command, its not width nor height but instead: %v",
					firstOpt.Name),
				s, i.Interaction)
			return
		}
	}

	err := discord.InteractionRespond(
		s,
		i.Interaction,
		dg.InteractionResponseChannelMessageWithSource,
		&dg.InteractionResponseData{
			Content: genAquarium(int(width), int(height)),
			// Embeds: []*dg.MessageEmbed{embed.MessageEmbed},
			// Files:  []*dg.File{{Name: "aquarium.png", Reader: aquariumR}, {Name: "boost.png", Reader: boostR}},
		},
	)
	if err != nil {
		discord.Error(fmt.Errorf("error responding to aquarium command interaction: %v", err), s, i.Interaction)
		return
	}
}
