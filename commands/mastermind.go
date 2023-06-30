package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/48thFlame/Boost-bot/data"
	"github.com/48thFlame/Boost-bot/discord"
	dg "github.com/bwmarrin/discordgo"
)

const (
	mastermindAPIUrl  = "http://35.238.73.100:8080/mastermind"
	mastermindGameLen = 7
)

func mewMastermindGame() (*MastermindGame, error) {
	res, err := http.Get(mastermindAPIUrl)
	if err != nil {
		return nil, err
	}

	game := &MastermindGame{}

	decoder := json.NewDecoder(res.Body)

	err = decoder.Decode(game)
	if err != nil {
		return nil, err
	}

	return game, nil
}

type MastermindGame struct {
	Won     bool     `json:"won"`
	Answer  [4]int   `json:"answer"`
	Guesses [][4]int `json:"guesses"`
	Results [][]int  `json:"results"`
}

func stringFromResultInt(r int) string {
	switch r {
	case 0:
		return "🔳"
	case 1:
		return "❎"
	case 2:
		return "✅"
	default:
		return ""
	}
}

func stringFromColorInt(c int) string {
	switch c {
	case 1:
		return "🟥"
	case 2:
		return "🟧"
	case 3:
		return "🟨"
	case 4:
		return "🟩"
	case 5:
		return "🟦"
	case 6:
		return "🟪"
	default:
		return "🔳"
	}
}

const (
	masterHighlighter    = "**"
	masterSecretEmoji    = "❓"
	masterBoardSeparator = " -- "
)

func (m *MastermindGame) gameToString() (str string) {
	str += masterHighlighter + "Answer:" + masterHighlighter + "\n"
	str += strings.Repeat(masterSecretEmoji+" ", 4) + masterBoardSeparator + strings.Repeat(stringFromResultInt(2)+" ", 4) + "\n"

	var guess [4]int
	var result []int

	for i := 0; i < mastermindGameLen; i++ {
		if len(m.Guesses) > i { //if guessed up until now, use the guess, otherwise use a blank/empty guess
			guess = m.Guesses[i]
			result = m.Results[i]
		} else {
			guess = [4]int{0, 0, 0, 0}
			result = []int{0, 0, 0, 0}
		}

		str += masterHighlighter + "Round " + fmt.Sprint(i+1, ":") + masterHighlighter
		str += "\n"

		for _, color := range guess {
			str += stringFromColorInt(color)
			str += " "
		}
		str += masterBoardSeparator
		for _, result := range result {
			str += stringFromResultInt(result)
			str += " "
		}
		if i != mastermindGameLen-1 { // if its not the last round then should add new line char
			str += "\n"
		}
	}

	return
}

func (game *MastermindGame) getAnswerString() string {
	return fmt.Sprintf(
		"%s %s %s %s",
		stringFromResultInt(game.Answer[0]),
		stringFromColorInt(game.Answer[1]),
		stringFromResultInt(game.Answer[2]),
		stringFromColorInt(game.Answer[3]))
}

type MastermindPostRequest struct {
	Game  *MastermindGame `json:"game"`
	Guess [4]int          `json:"guess"`
}

func getMasterColorIntFromMastermindCommandOptionsStringValue(s string) int {
	switch s {
	case "🟥 - red":
		return 1
	case "🟧 - orange":
		return 2
	case "🟨 - yellow":
		return 3
	case "🟩 - green":
		return 4
	case "🟦 - blue":
		return 5
	case "🟪 - purple":
		return 6
	default:
		return 0
	}
}

func getMastermindGame(id string) (*MastermindGame, error) {
	var err error
	hasGame := data.DataExists(data.GetMastermindFileName(id))
	game := &MastermindGame{}

	if hasGame {
		err = data.LoadData(data.GetMastermindFileName(id), game)
		if err != nil {
			return nil, err
		}
	} else {
		game, err = mewMastermindGame()
		if err != nil {
			return nil, err
		}
	}
	return game, nil
}

func Mastermind(s *dg.Session, i *dg.InteractionCreate) {
	// init vars
	subCommands := i.ApplicationCommandData().Options
	cmdName := subCommands[0].Name
	id := discord.GetInteractionUser(i.Interaction).ID
	interaction := i.Interaction

	switch cmdName {
	case "rules":
		// responding in python ;)
	case "new-game":
		mastermindNewGame(s, interaction, id)
	case "main":
		fallthrough
	case "guess":
		// mastermindMainOrGuess()
		mastermindMainOrGuess(s, interaction, subCommands[0].Options, id, cmdName == "guess")
	default:
		discord.Error(fmt.Errorf("undefined mastermind subcommand: %v", cmdName), s, interaction)
		return
	}
}

func mastermindNewGame(s *dg.Session, interaction *dg.Interaction, id string) {
	hasGame := data.DataExists(data.GetMastermindFileName(id))

	// only if hasGame needs to start a new-game because will make new otherwise when calling other command
	if hasGame {
		err := data.DeleteData(data.GetMastermindFileName(id))
		if err != nil {
			discord.Error(fmt.Errorf("error deleting mastermind game: %v", err), s, interaction)
			return
		}

		user, err := data.LoadUser(id)
		if err != nil {
			discord.Error(fmt.Errorf("error loading user: %v", err), s, interaction)
			return
		}

		user.Stats.Mastermind.Losses++
		err = data.SaveData(data.GetUserFileName(id), user)
		if err != nil {
			discord.Error(fmt.Errorf("error saving user: %v", err), s, interaction)
			return
		}
	}

	discord.InteractionRespond(
		s,
		interaction,
		discord.InstaMessage,
		&dg.InteractionResponseData{
			Content: "New mastermind game successfully started!",
		},
	)
}

func (game *MastermindGame) makeGuess(guess [4]int) (*MastermindGame, error) {
	postData := MastermindPostRequest{
		Game:  game,
		Guess: guess,
	}

	bodyData, err := json.Marshal(postData)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(mastermindAPIUrl, "application/json", bytes.NewReader(bodyData))
	if err != nil {
		return nil, err
	}

	jsonDecoder := json.NewDecoder(resp.Body)

	err = jsonDecoder.Decode(game)
	if err != nil {
		return nil, err
	}

	return game, nil
}

func mastermindMainOrGuess(
	s *dg.Session,
	interaction *dg.Interaction,
	options []*dg.ApplicationCommandInteractionDataOption,
	id string,
	guessCmd bool) {

	// init vars
	var won, lost bool
	var err error

	game, err := getMastermindGame(id)
	if err != nil {
		discord.Error(fmt.Errorf("error getting/loading mastermind game: %v", err), s, interaction)
		return
	}

	// if guessed should guess otherwise only needs to display
	if guessCmd {
		guess := [4]int{}

		for cI, c := range options {
			cStr := c.StringValue()
			guess[cI] = getMasterColorIntFromMastermindCommandOptionsStringValue(cStr)
		}

		game, err = game.makeGuess(guess)
		if err != nil {
			discord.Error(fmt.Errorf("error guessing on mastermind game: %v", err), s, interaction)
			return
		}

		won = game.Won
		lost = len(game.Guesses) >= mastermindGameLen && !won
	}

	err = data.SaveData(data.GetMastermindFileName(id), game)
	if err != nil {
		discord.Error(fmt.Errorf("error saving mastermind game: %v", err), s, interaction)
	}

	embed := discord.NewEmbed().
		SetupEmbed().
		SetAuthor("attachment://mastermind.png", "Mastermind", "").
		SetDescription(game.gameToString())
	if won {
		embed.SetTitle("Congratulations! 🥳 You won!")
	} else if lost {
		embed.SetTitle(fmt.Sprintf("You lost! 😢\nThe answer was: %v", game.getAnswerString()))
	}

	// should delete data if won or lost
	if won || lost {
		// should delete the mastermind game
		data.DeleteData(data.GetMastermindFileName(id))

		// should update the user stats
		user, err := data.LoadUser(id)
		if err != nil {
			discord.Error(fmt.Errorf("error loading user: %v", err), s, interaction)
		}
		if won {
			user.Stats.Mastermind.Wins++
			user.Stats.Mastermind.Rounds += len(game.Guesses)
		} else {
			user.Stats.Mastermind.Losses++
		}

		// should save the user
		err = data.SaveData(data.GetUserFileName(id), user)
		if err != nil {
			discord.Error(fmt.Errorf("error saving user: %v", err), s, interaction)
		}
	}

	boostR, err := os.Open("./assets/boost.png")
	if err != nil {
		discord.Error(fmt.Errorf("error opening boost.png: %v", err), s, interaction)
	}

	mastermindR, err := os.Open("./assets/mastermind.png")
	if err != nil {
		discord.Error(fmt.Errorf("error opening mastermind.png: %v", err), s, interaction)
	}

	err = discord.InteractionRespond(
		s,
		interaction,
		discord.InstaMessage,
		&dg.InteractionResponseData{
			Embeds: []*dg.MessageEmbed{embed.MessageEmbed},
			Files:  []*dg.File{{Name: "boost.png", Reader: boostR}, {Name: "mastermind.png", Reader: mastermindR}},
		},
	)
	if err != nil {
		discord.Error(fmt.Errorf("error responding to mastermind command interaction: %v", err), s, interaction)
	}
}
