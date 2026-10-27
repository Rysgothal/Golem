package commands

import (
	"github.com/bwmarrin/discordgo"
)

type Command struct {
	Definition *discordgo.ApplicationCommand
	Handler    func(session *discordgo.Session, interaction *discordgo.InteractionCreate)
}

func GetCommands() []Command {
	return []Command{pingCmd}
}
