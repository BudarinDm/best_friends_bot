package config

import (
	"errors"
	"os"
)

type Config struct {
	Bot Bot
}

type Bot struct {
	Token string
	Debug string
}

func ReadConfig() (*Config, error) {

	var config Config
	var err error

	//app parse

	config.Bot.Debug = os.Getenv("DEBUG")
	if config.Bot.Debug == "true" {
		config.Bot.Token = os.Getenv("BOT_BF_TOKEN_STAGE")
		if config.Bot.Token == "" {
			return nil, errors.New("error BOT_BF_TOKEN_STAGE")
		}
	} else {
		config.Bot.Token = os.Getenv("BOT_BF_TOKEN")
		if config.Bot.Token == "" {
			return nil, errors.New("error BOT_BF_TOKEN")
		}
	}

	return &config, err

}
