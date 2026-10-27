package discord

import (
	// "log"

	"fmt"
	"golem/internal/discord/commands"
	"log"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	session  *discordgo.Session
	commands map[string]commands.Command
	guildID  string
}

func New(session *discordgo.Session, cmds []commands.Command) *Bot {
	dict := make(map[string]commands.Command, len(cmds))

	for _, command := range cmds {
		dict[command.Definition.Name] = command
	}

	return &Bot{session: session, commands: dict}
}

// Before .Open()
func (b *Bot) Assign() {
	b.session.AddHandler(b.getGuildID)
	b.session.AddHandler(b.route)
}

// After .Open()
func (b *Bot) Register() error {
	cmds := make([]*discordgo.ApplicationCommand, 0, len(b.commands))
	for _, c := range b.commands {
		cmds = append(cmds, c.Definition)
	}
	if b.session.State == nil || b.session.State.User == nil {
		return fmt.Errorf("bot ainda não recebeu o Ready")
	}
	_, err := b.session.ApplicationCommandBulkOverwrite(
		b.session.State.User.ID, b.guildID, cmds,
	)
	return err
}

// Unknow command Handler
func (b *Bot) route(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("command panic: %v", r)
		}
	}()

	if interaction.Type != discordgo.InteractionApplicationCommand {
		return
	}

	name := interaction.ApplicationCommandData().Name
	cmd, ok := b.commands[name]

	if !ok {
		log.Printf("commando unknow: %q", name)
		return
	}

	// TODO: checagem de admin aqui

	cmd.Handler(session, interaction)
}

func (b *Bot) getGuildID(session *discordgo.Session, guild *discordgo.GuildCreate) {
	b.guildID = guild.ID
}