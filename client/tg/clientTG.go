package tg

import (
	"compendium/config"
	"compendium/models"
	"compendium/storage"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mentalisit/logger"
)

type Telegram struct {
	ChanMessage chan models.IncomingMessage
	t           *tgbotapi.BotAPI
	log         *logger.Logger
	storage     *storage.Storage
}

func NewTelegram(log *logger.Logger, cfg *config.ConfigBot, st *storage.Storage) *Telegram {
	tgBot, err := tgbotapi.NewBotAPI(cfg.Token.TokenTelegram)
	if err != nil {
		log.Panic("ошибка подключения к телеграм " + err.Error())
	}

	tgBot.Debug = false
	fmt.Printf("Бот TELEGRAM загружен  %s\n", tgBot.Self.UserName)

	tg := &Telegram{
		ChanMessage: make(chan models.IncomingMessage, 10),
		t:           tgBot,
		log:         log,
		storage:     st,
	}
	go tg.update()
	return tg
}
