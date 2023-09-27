package service

import (
	"bytes"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
	"math/rand"
	"net/http"
	"time"
)

func (s Bot) SendPhotoIsWord(update tgbotapi.Update, photoName string, countPhoto int) error {
	randSource := rand.NewSource(time.Now().UnixNano())
	randObj := rand.New(randSource)
	randVasya := randObj.Intn(countPhoto)

	log.Print(randVasya)

	url := fmt.Sprintf("https://s3.timeweb.com/3c0377c1-core/vasya/%s-%d.jpeg", photoName, randVasya)
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error http.Get(url) %s", update.Message.From.UserName)
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(res.Body)
	if err != nil {
		return fmt.Errorf("error buf.ReadFrom(res.Body) %s", update.Message.From.UserName)
	}

	file := tgbotapi.FileBytes{
		Name:  "",
		Bytes: buf.Bytes(),
	}
	photo := tgbotapi.NewPhoto(update.Message.Chat.ID, file)
	//msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	photo.ReplyToMessageID = update.Message.MessageID

	err = res.Body.Close()
	if err != nil {
		return fmt.Errorf("error res.Body.Close() %s", update.Message.From.UserName)
	}

	_, err = s.Bot.Send(photo)
	if err != nil {
		return fmt.Errorf("error bot.Send(photo) %s", update.Message.From.UserName)
	}

	return nil
}
