package config

import (
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DiscordToken string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil && !errors.Is(err, fs.ErrNotExist) {
		return nil, fmt.Errorf("reading .env: %w", err)
	}

	token := os.Getenv("BOT_DISCORD_TOKEN")
	if token == "" {
		return nil, errors.New("BOT_DISCORD_TOKEN not defined: create .env from .env.example")
	}

	return &Config{DiscordToken: token}, nil
}
