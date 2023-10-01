package service

import (
	"best_friends_bot/pkg/logger"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

func (a *App) startConsumers(ctx context.Context) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := a.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			logger.Infof("userName- %s , userID- %d , message- %s, data- %v", update.Message.From.UserName, update.Message.From.ID, update.Message.Text, update.Message.From)

			err := a.messageByTrigger(ctx, update)
			if err != nil {
				logger.Errorf("updateChan: %s", err.Error())
				continue
			}
		}
	}
}

func (a *App) messageByTrigger(ctx context.Context, update tgbotapi.Update) error {
	triggers, err := a.logic.GetTriggers(ctx)
	if err != nil {
		return fmt.Errorf("error SendPhotoIsWord: %s", err.Error())
	}

	for _, t := range triggers {
		if strings.Contains(strings.ToLower(update.Message.Text), t.Trigger) {
			err = a.logic.SendMessageIsWord(ctx, update, t.Trigger)
			if err != nil {
				return fmt.Errorf("error SendMessageIsWord: %s", err.Error())
			}
		}
	}
	return nil
}
