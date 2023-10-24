package main

import (
	"github.com/gofor-little/env"
	bot "github.com/timhi/GooodReadsBot/Bot"
	backend "github.com/timhi/GooodReadsBot/backend"
)

func main() {
	if err := env.Load(".env"); err != nil {
		panic(err)
	}
	token, err := env.MustGet("DISCORD_TOKEN")
	port, err := env.MustGet("BACKEND_PORT")
	if err != nil {
		panic(err)
	}
	backendChannel := make(chan bool)
	botChannel := make(chan bool)
	go backend.Start(port, backendChannel)
	go bot.Start(token, botChannel)
	<-backendChannel
	<-botChannel
}
