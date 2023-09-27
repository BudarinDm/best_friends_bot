package main

import (
	"best_friends_bot/internal/config"
	"best_friends_bot/pkg/logger"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

func main() {
	logger.Init()

	cfg, err := config.ReadConfig()
	if err != nil {
		logger.Fatalf("error read config: %v", err)
	}

	bot, err := tgbotapi.NewBotAPI(cfg.Bot.Token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			message := update.Message.Text
			isOzon := strings.Contains(strings.ToLower(message), "озон")

			fmt.Println(isOzon)
			//isOzon := strings.Contains(message, "озон")

			if isOzon {
				// If we got a message
				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

				randSource := rand.NewSource(time.Now().UnixNano())
				randObj := rand.New(randSource)
				randVasya := randObj.Intn(3)

				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath(fmt.Sprintf("img/vasya%d.jpeg", randVasya)))
				//msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				photo.ReplyToMessageID = update.Message.MessageID

				_, err := bot.Send(photo)
				if err != nil {
					fmt.Println(err)
					return
				}
			}

		}
	}
}

func createWordRegex(word string) *regexp.Regexp {
	return regexp.MustCompile(`\b` + regexp.QuoteMeta(word) + `\b`)
}
