package service

import (
	"best_friends_bot/internal/config"
	"best_friends_bot/pkg/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	Cfg *config.Config
	Bot *tgbotapi.BotAPI
}

func NewBot(cfg *config.Config) Bot {
	bot, err := tgbotapi.NewBotAPI(cfg.Bot.Token)
	if err != nil {
		logger.Fatalf("error bot create %s", err.Error())
	}
	logger.Infof("Authorized on account %s", bot.Self.UserName)
	return Bot{
		Cfg: cfg,
		Bot: bot,
	}
}
