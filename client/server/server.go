package server

import (
	"compendium/client/ds"
	"compendium/config"
	"compendium/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mentalisit/logger"
	"os"
)

type Server struct {
	log *logger.Logger
	db  *storage.Storage
	ds  *ds.Discord
}

func NewServer(log *logger.Logger, cfg *config.ConfigBot, st *storage.Storage, d *ds.Discord) *Server {
	s := &Server{
		log: log,
		db:  st,
		ds:  d,
	}
	fmt.Println("Server load")
	go s.RunServer(cfg.Port)
	return s
}

func (s *Server) RunServer(port string) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// Обработчик для принятия сообщений от DiscordService
	router.OPTIONS("/compendium/applink/identities", s.Check)
	router.GET("/compendium/applink/identities", s.CheckIdentityHandler)

	router.OPTIONS("/compendium/applink/connect", s.Check)
	router.POST("/compendium/applink/connect", s.CheckConnectHandler)

	router.OPTIONS("/compendium/cmd/syncTech/:mode", s.Check)
	router.POST("/compendium/cmd/syncTech/:mode", s.CheckSyncTechHandler)

	router.OPTIONS("/compendium/cmd/corpdata", s.Check)
	router.GET("/compendium/cmd/corpdata", s.CheckCorpDataHandler)

	router.OPTIONS("/compendium/applink/refresh", s.Check)
	router.GET("/compendium/applink/refresh", s.CheckRefreshHandler)
	fmt.Println("Running port:" + port)
	err := router.Run(":" + port)
	if err != nil {
		s.log.ErrorErr(err)
		os.Exit(1)
	}
}
