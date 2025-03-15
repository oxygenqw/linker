package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Telegram Telegram `yaml:"TELEGRAM"`
	Server   Server   `yaml:"SERVER"`
}

type Telegram struct {
	BotToken string `yaml:"BOT_TOKEN" env:"BOT_TOKEN" env-required:"true"`
	AppURL   string `yaml:"APP_URL"   env:"APP_URL"   env-required:"true"`
}

type Server struct {
	Host string `yaml:"HOST" env:"HOST"`
	Port string `yaml:"PORT" env:"PORT" env-required:"true"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "config.yaml"
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	return &cfg
}
