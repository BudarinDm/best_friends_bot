package logic

import (
	"best_friends_bot/internal/model"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"math/rand"
	"time"
)

func (l Logic) SendDRNextCommand(ctx context.Context, update tgbotapi.Update) error {
	t := time.Now()

	bd, err := l.repo.GetBDDateNext(ctx, t)
	if err != nil {
		return fmt.Errorf("error GetBDDates: %s", err.Error())
	}
	msgText := fmt.Sprintf(`Ближайший др у %s. %d %s`, bd.FIO, bd.Day, bd.MonthText)

	newMsg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
	newMsg.ReplyToMessageID = update.Message.MessageID

	_, err = l.bot.Send(newMsg)
	if err != nil {
		return fmt.Errorf("error bot.Send(msg): %s", err.Error())
	}

	return nil
}

func (l Logic) SendBadRequest(ctx context.Context, update tgbotapi.Update, command string) error {
	msgText := fmt.Sprintf(
		"Хватит вводить всякую хуйню.\nДоступные команды для %s:\n%s", command, model.MessageInfoDr,
	)

	newMsg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
	newMsg.ReplyToMessageID = update.Message.MessageID

	_, err := l.bot.Send(newMsg)
	if err != nil {
		return fmt.Errorf("error bot.Send(msg): %s", err.Error())
	}

	return nil
}

func (l Logic) SendDRCommand(ctx context.Context, update tgbotapi.Update) error {
	bd, err := l.repo.GetBDDates(ctx)
	if err != nil {
		return fmt.Errorf("error GetBDDates: %s", err.Error())
	}
	msgText := "Дни рождения:\n"

	for i, b := range bd {
		msgText += fmt.Sprintf("%d. %s: %d %s.\n", i+1, b.FIO, b.Day, b.MonthText)
	}

	newMsg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
	newMsg.ReplyToMessageID = update.Message.MessageID

	_, err = l.bot.Send(newMsg)
	if err != nil {
		return fmt.Errorf("error bot.Send(msg): %s", err.Error())
	}

	return nil
}

func (l Logic) SendMessage(ctx context.Context, update tgbotapi.Update, message string) error {

	newMsg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	newMsg.ReplyToMessageID = update.Message.MessageID

	_, err := l.bot.Send(newMsg)
	if err != nil {
		return fmt.Errorf("error bot.Send(msg): %s", err.Error())
	}

	return nil
}

func (l Logic) SendMessageIsWord(ctx context.Context, update tgbotapi.Update, trigger string) error {
	messages, err := l.repo.GetMessageByTrigger(ctx, trigger)
	if err != nil {
		return fmt.Errorf("error GetMessageByTrigger: %s", err.Error())
	}

	randSource := rand.NewSource(time.Now().UnixNano())
	randObj := rand.New(randSource)
	randIndex := randObj.Intn(len(messages))
	text := messages[randIndex]

	newMsg := tgbotapi.NewMessage(update.Message.Chat.ID, text.Message)
	newMsg.ReplyToMessageID = update.Message.MessageID

	_, err = l.bot.Send(newMsg)
	if err != nil {
		return fmt.Errorf("error bot.Send(msg): %s", err.Error())
	}

	return nil
}

func (l Logic) SendPhotoIsWord(update tgbotapi.Update, photoUrl string) error {
	photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath(photoUrl))
	photo.ReplyToMessageID = update.Message.MessageID

	_, err := l.bot.Send(photo)
	if err != nil {
		return fmt.Errorf("error bot.Send(photo) %s", update.Message.From.UserName)
	}

	return nil
}
