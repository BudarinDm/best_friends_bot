package service

import (
	"best_friends_bot/internal/config"
	"best_friends_bot/internal/logic"
	"best_friends_bot/pkg/logger"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"math/rand"
	"net/http"

	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// App основная структура для приложения
type App struct {
	config *config.Config
	logic  *logic.Logic
	router *gin.Engine
	bot    *tgbotapi.BotAPI
}

func NewApp(config *config.Config, logic *logic.Logic, router *gin.Engine, bot *tgbotapi.BotAPI) *App {
	rand.NewSource(time.Now().UTC().UnixNano())

	return &App{
		config: config,
		logic:  logic,
		router: router,
		bot:    bot,
	}
}

func (a *App) Start() error {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)
	go a.startConsumers(context.Background())
	go func() {
		sig := <-sigs
		logger.Infof("sig : %s", sig)
		logger.Infof("stoping subscriptions...")
		done <- true
	}()

	<-done
	return nil
}

func (a *App) StartServer(port string) {
	a.router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, ResponseType, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", port),
		Handler:        a.router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	logger.Infof("Сервис запущен на порту %v", port)
	err := s.ListenAndServe()
	if err != nil {
		logger.Fatalf("ошибка в работе сервиса error: %s", err.Error())
	}
}
