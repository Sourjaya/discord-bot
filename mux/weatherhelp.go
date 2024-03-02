package mux

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

func (m *Mux) WeatherHelp(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	ds.ChannelMessageSend(dm.ChannelID, "Ping = "+(time.Since(dm.Timestamp)).String())
}
