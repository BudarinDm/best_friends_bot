package config

import (
	"errors"
	"os"
)

type Config struct {
	App AppConfig
	Bot Bot
	DB  DBConfig
}

type Bot struct {
	Token string
	Debug string
}

type DBConfig struct {
	DBString string
}

type AppConfig struct {
	Port     string
	LogLevel string
}

func ReadConfig() (*Config, error) {

	var config Config
	var err error

	//app parse
	config.App.LogLevel = os.Getenv("LOG_LEVEL")
	if config.App.LogLevel == "" {
		config.App.LogLevel = "debug"
	}
	config.App.Port = os.Getenv("SERVER_PORT")
	if config.App.Port == "" {
		config.App.Port = "80"
	}

	//bot parse

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

	//db parse

	config.DB.DBString = os.Getenv("DBSTRING")
	if config.DB.DBString == "" {
		return nil, errors.New("not specified DBSTRING")
	}

	return &config, err

}
