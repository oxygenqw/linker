package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Telegram Telegram `yaml:"TELEGRAM"`
	Server   Server   `yaml:"SERVER"`
	Database Database `yaml:"DATABASE"`
}

type Telegram struct {
	BotToken string `yaml:"BOT_TOKEN" env:"BOT_TOKEN" env-required:"true"`
	AppURL   string `yaml:"APP_URL"   env:"APP_URL"   env-required:"true"`
}

type Server struct {
	Host string `yaml:"HOST" env:"HOST"`
	Port string `yaml:"PORT" env:"PORT" env-required:"true"`
}

type Database struct {
	Host     string `yaml:"HOST"     env:"DB_HOST"     env-required:"true"`
	Port     string `yaml:"PORT"     env:"DB_PORT"     env-required:"true"`
	User     string `yaml:"USER"     env:"DB_USER"     env-required:"true"`
	Password string `yaml:"PASSWORD" env:"DB_PASSWORD" env-required:"true"`
	Name     string `yaml:"NAME"     env:"DB_NAME"     env-required:"true"`
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
