package ds

import (
	"compendium/config"
	"compendium/models"
	"compendium/storage"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/mentalisit/logger"
)

type Discord struct {
	s           *discordgo.Session
	log         *logger.Logger
	storage     *storage.Storage
	ChanMessage chan models.IncomingMessage
}

func NewDiscord(log *logger.Logger, st *storage.Storage, cfg *config.ConfigBot) *Discord {
	ds, err := discordgo.New("Bot " + cfg.Token.TokenDiscord)
	if err != nil {
		log.Panic("Ошибка запуска дискорда" + err.Error())
		return nil
	}

	err = ds.Open()
	if err != nil {
		log.Panic("Ошибка открытия ДС " + err.Error())
	}

	DS := &Discord{
		s:           ds,
		log:         log,
		storage:     st,
		ChanMessage: make(chan models.IncomingMessage, 20),
	}

	ds.AddHandler(DS.messageHandler)
	//ds.AddHandler(DS.messageReactionAdd)

	fmt.Println("Бот Дискорд загружен ")
	return DS
}
