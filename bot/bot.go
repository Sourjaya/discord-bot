package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Sourjaya/discord-bot/mux"
	"github.com/bwmarrin/discordgo"
)

var (
	BotToken string
	Router   *mux.Mux
)

func Start() {
	// Create a new Discord session using the provided bot token.
	discord, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatalf("error creating Discord session: %v", err)
		return
	}
	Router = mux.New()
	Router.Prefix = "!"
	// Register the mSend func as a callback for MessageCreate events.
	discord.AddHandler(Router.OnMessageCreate)
	Router.Route("ping", "Ping command that returns latency", Router.Ping)
	Router.Route("weather", "Weather command that returns  ", Router.Weather)
	Router.Route("weatherhelp", "Weather command that returns  ", Router.WeatherHelp)
	// Open a websocket connection to Discord and begin listening.
	err = discord.Open()
	if err != nil {
		log.Fatalf("error opening websocket connection: %v", err)
		return
	}
	defer discord.Close()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running")

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-s
}
