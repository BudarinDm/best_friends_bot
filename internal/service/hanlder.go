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
			logger.Infof("userName- %s , userID- %d , message- %s", update.Message.From.UserName, update.Message.From.ID, update.Message.Text)

			if update.Message.IsCommand() {
				err := a.commandHandler(ctx, update)
				if err != nil {
					logger.Errorf("commandHandler: %s", err.Error())
					continue
				}
				continue
			}

			err := a.messageByTrigger(ctx, update)
			if err != nil {
				logger.Errorf("messageByTrigger: %s", err.Error())
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

func (a *App) commandHandler(ctx context.Context, update tgbotapi.Update) error {
	command := update.Message.Command()

	if strings.ToLower(command) == "dr" {
		arg := update.Message.CommandArguments()
		switch arg {
		case "":
			err := a.logic.SendDRCommand(ctx, update)
			if err != nil {
				return fmt.Errorf("error SendDRCommand: %s", err.Error())
			}
		case "next":
			err := a.logic.SendDRNextCommand(ctx, update)
			if err != nil {
				return fmt.Errorf("error SendDRNextCommand: %s", err.Error())
			}
		default:
			err := a.logic.SendBadRequest(ctx, update, command)
			if err != nil {
				return fmt.Errorf("error SendBadRequest: %s", err.Error())
			}
		}
	}

	return nil
}
