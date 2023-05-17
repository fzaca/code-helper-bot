package config

import (
	"github.com/Xukay101/code-helper-bot/src/utils"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Bot struct {
		Token  string `env:"DISCORD_TOKEN"`
		Prefix string `env-default:"%"`
	}
	Database struct {
		Port     string `env:"DB_PORT"`
		Host     string `env:"DB_HOST"`
		Name     string `env:"DB_NAME"`
		User     string `env:"DB_USER"`
		Password string `env:"DB_PASSWORD"`
	}
}

func GetConfig() Config {
	var config Config
	err := cleanenv.ReadEnv(&config)
	utils.FatalOnError("Error getting settings", err)
	return config
}
