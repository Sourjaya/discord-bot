package mux

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	WeatherAPIToken string
)

func (m *Mux) Weather(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	cWeather := weather.getWeather(dm.Content, WeatherAPIToken)
	if _, err := ds.ChannelMessageSendComplex(dm.ChannelID, cWeather); err != nil {
		log.Fatalf("Error while sending message over channel: %v", err)
	}
}
