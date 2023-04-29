package discord

import (
	"log"
	"os"
	"os/exec"

	dg "github.com/bwmarrin/discordgo"
)

type SlashCommandHandlerType func(s *dg.Session, i *dg.InteractionCreate)

func NewBot(tokenFilePath, pyInterpreter, pyCommandsFile string, commands map[string]SlashCommandHandlerType) (*Bot, error) {
	bot := &Bot{}

	token, err := os.ReadFile(tokenFilePath)
	if err != nil {
		return nil, err
	}

	s, err := dg.New("Bot " + string(token))
	if err != nil {
		return nil, err
	}

	s.AddHandler(func(s *dg.Session, i *dg.InteractionCreate) {
		if i.Type == dg.InteractionApplicationCommand {
			if cmdHandler, exists := bot.cmdsHandlers[i.ApplicationCommandData().Name]; exists {
				cmdHandler(s, i)
			} else {
				log.Printf("Command '%v' was executed by user, bot no handler was defined.\n", i.ApplicationCommandData().Name)
			}
		}
	})

	s.AddHandler(func(s *dg.Session, r *dg.Ready) {
		log.Printf("%v is now online!\n", bot.S.State.User)
	})
	bot.cmdsHandlers = make(map[string]SlashCommandHandlerType)

	for name, handler := range commands {
		bot.addCommandHandler(name, handler)
	}

	bot.S = s
	bot.pyInterpreter = pyInterpreter
	bot.pyCommandsFile = pyCommandsFile

	return bot, nil
}

type Bot struct {
	S                             *dg.Session
	cmdsHandlers                  map[string]SlashCommandHandlerType
	pyInterpreter, pyCommandsFile string
}

func (b *Bot) Start() error {
	go func() {
		err := b.runPyScript()
		if err != nil {
			log.Fatalf("Error while running python script: %v\n", err)
		}
	}()

	return b.S.Open()
}

func (b *Bot) runPyScript() (err error) {
	cmd := exec.Command(b.pyInterpreter, "-OO", b.pyCommandsFile)
	cmd.Stdout = log.Default().Writer()

	err = cmd.Run()

	return
}

func (b *Bot) addCommandHandler(name string, handler SlashCommandHandlerType) {
	b.cmdsHandlers[name] = handler
}
