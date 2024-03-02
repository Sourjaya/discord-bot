package main

import (
	"flag"

	"github.com/Sourjaya/discord-bot/bot"
	"github.com/Sourjaya/discord-bot/mux"
)

// Variables used for command line parameters
var (
	BotToken        string
	WeatherAPIToken string
)

func init() {
	flag.StringVar(&BotToken, "b", "", "Bot Token")
	flag.StringVar(&WeatherAPIToken, "w", "", "Weather API Token")
	flag.Parse()
}

func main() {
	// Save Bot and API keys and start bot
	bot.BotToken = BotToken
	mux.WeatherAPIToken = WeatherAPIToken
	bot.Start()
}
