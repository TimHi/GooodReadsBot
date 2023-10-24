package bot

import (
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/charmbracelet/log"
	"github.com/timhi/GooodReadsBot/Bot/cmd"
)

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name: "basic-command",
			// All commands and options must have a description
			// Commands/options without description will fail the registration
			// of the command.
			Description: "Basic command",
		},
	}
)

func Start(token string, ch chan<- bool) {
	cmd.SearchBook("Aal")
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	err = s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	s.ApplicationCommandBulkOverwrite(s.State.User.ID, "", commands)
	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	log.Info("Press Ctrl+C to exit")
	<-stop

	log.Info("Gracefully shutting down.")
	ch <- true
}
