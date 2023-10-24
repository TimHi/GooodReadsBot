package bot

import (
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/charmbracelet/log"
	"github.com/timhi/GooodReadsBot/Bot/cmd"
)

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "book",
			Description: "Search a book on GoodReads",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "query",
					Description: "Enter your book title",
					Required:    true,
				},
			},
		},
	}
)

func Start(token string, ch chan<- bool) {

	var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"book": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options

			bookResult, err := cmd.SearchBook(options[0].StringValue())

			if err != nil {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: err.Error(),
					}})
			} else {
				log.Info(bookResult)
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Embeds: []*discordgo.MessageEmbed{
							{
								Title:       bookResult.Title,
								Description: strings.Join(bookResult.Authors, " - "),
								URL:         "https://goodreads.com/" + bookResult.Link,
								Thumbnail: &discordgo.MessageEmbedThumbnail{
									URL: bookResult.Cover,
								},
								Fields: []*discordgo.MessageEmbedField{
									{
										Name:  "Ratings",
										Value: bookResult.Rating,
									},
								},
							},
						},
					},
				})
			}
		},
	}

	s, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

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
