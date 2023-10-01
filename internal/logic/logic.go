package logic

import (
	"best_friends_bot/internal/config"
	db "best_friends_bot/internal/repository"
	"best_friends_bot/pkg/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Logic struct {
	Cfg  *config.Config
	bot  *tgbotapi.BotAPI
	repo *db.Repo
}

func NewLogic(cfg *config.Config, repo *db.Repo, bot *tgbotapi.BotAPI) *Logic {
	logger.Infof("Authorized on account %s", bot.Self.UserName)
	return &Logic{
		Cfg:  cfg,
		bot:  bot,
		repo: repo,
	}
}
