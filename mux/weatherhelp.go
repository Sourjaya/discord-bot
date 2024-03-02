package mux

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func (m *Mux) WeatherHelp(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	helpMessage := &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{{
			Type:        discordgo.EmbedTypeRich,
			Title:       "Help",
			Description: "Help for Weather Bot",
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Command Syntax",
					Value:  "!weather <location name>",
					Inline: true,
				},
				{
					Name:   "Command Example 1",
					Value:  "!weather kolkata",
					Inline: true,
				},
				{
					Name:   "Command Example 2",
					Value:  "!weather Rio De Janeiro",
					Inline: true,
				},
			},
		},
		},
	}
	if _, err := ds.ChannelMessageSendComplex(dm.ChannelID, helpMessage); err != nil {
		log.Fatalf("Error while sending message over channel: %v", err)
	}
}
