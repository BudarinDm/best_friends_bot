package main

import (
	"best_friends_bot/internal/config"
	"best_friends_bot/internal/model"
	"best_friends_bot/internal/service"
	"best_friends_bot/pkg/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

func main() {
	logger.Init()

	cfg, err := config.ReadConfig()
	if err != nil {
		logger.Fatalf("error read config: %v", err)
	}

	srv := service.NewBot(cfg)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := srv.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			message := update.Message.Text

			if isVasya(message) {
				err = srv.SendPhotoIsWord(update, model.PhotoVasyaName, 3)
				if err != nil {
					logger.Errorf("error SendPhotoIsWord: %v", err)
					continue
				}
			}

			if isOzon(message) {
				err = srv.SendPhotoIsWord(update, model.PhotoOzonName, 8)
				if err != nil {
					logger.Errorf("error SendPhotoIsWord: %v", err)
					continue
				}
			}

		}
	}
}

func isVasya(message string) bool {
	if strings.Contains(strings.ToLower(message), "вася") {
		return true
	}
	if strings.Contains(strings.ToLower(message), "васек") {
		return true
	}
	if strings.Contains(strings.ToLower(message), "василий") {
		return true
	}
	return false
}

func isOzon(message string) bool {
	if strings.Contains(strings.ToLower(message), "озон") {
		return true
	}
	if strings.Contains(strings.ToLower(message), "работа") {
		return true
	}
	if strings.Contains(strings.ToLower(message), "ozon") {
		return true
	}
	return false
}
