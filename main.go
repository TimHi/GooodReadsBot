package main

import (
	"github.com/gofor-little/env"
	bot "github.com/timhi/GooodReadsBot/bot"
)

func main() {
	if err := env.Load(".env"); err != nil {
		panic(err)
	}

	token, err := env.MustGet("DISCORD_TOKEN")
	if err != nil {
		panic(err)
	}

	botChannel := make(chan bool)
	go bot.Start(token, botChannel)
	<-botChannel
}
