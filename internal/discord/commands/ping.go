package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var pingCmd = Command{
	Definition: &discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Verifica se o bot está respondendo",
	},

	Handler: func(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
		err := session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Pong",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})

		if err != nil {
			log.Printf("ping: respondendo: %v", err)
		}
	},
}
