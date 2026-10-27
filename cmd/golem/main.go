package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"golem/internal/config"
	"golem/internal/discord"
	"golem/internal/discord/commands"
)

func main() {
	fmt.Println("loading config...")
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("establishing connection...")
	dg, err := discordgo.New("Bot " + cfg.DiscordToken)
	if err != nil {
		log.Fatal(err)
	}
	dg.Identify.Intents = discordgo.IntentsGuilds

	fmt.Println("turning on bot...")
	bot := discord.New(dg, commands.GetCommands())
	bot.Assign()

	if err := dg.Open(); err != nil {
		log.Fatalf("Conecting Discord: %v", err)
	}
	defer dg.Close()

	fmt.Println("applying config on bot...")
	if err := bot.Register(); err != nil {
		log.Fatalf("Bot functions: %v", err)
	}

	log.Println("Bot initialized")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
}
