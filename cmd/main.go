package main

import (
	"best_friends_bot/internal/config"
	"best_friends_bot/internal/logic"
	db "best_friends_bot/internal/repository"
	"best_friends_bot/internal/service"
	"best_friends_bot/pkg/logger"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	logger.Init()

	cfg, err := config.ReadConfig()
	if err != nil {
		logger.Fatalf("error read config: %v", err)
	}

	bot, err := tgbotapi.NewBotAPI(cfg.Bot.Token)
	if err != nil {
		logger.Fatalf("error bot create %s", err.Error())
	}

	repo, err := db.NewRepo(&cfg.DB)
	if err != nil {
		logger.Fatalf("Failed to initialize DB: %s", err)
	}
	defer repo.Close()

	lgc := logic.NewLogic(cfg, repo, bot)

	router := gin.Default()
	srv := service.NewApp(cfg, lgc, router, bot) // service использует logic для обработки событий

	go srv.StartServer(cfg.App.Port)

	if err = srv.Start(); err != nil {
		logger.Fatalf("Ошибка при работе сервиса: %s", err)
	}
}
