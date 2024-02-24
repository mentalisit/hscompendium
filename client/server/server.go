package server

import (
	"compendium/models"
	"compendium/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mentalisit/logger"
	"os"
)

type Server struct {
	log *logger.Logger
	db  db
}
type db interface {
	ReadIdentity() []models.Identity
	ReadIdentityByToken(token string) models.Identity
	InsertIdentity(c models.Identity)
	UpdateIdentity(c models.Identity)
	ReadSyncData(token string) models.SyncData
}

func NewServer(log *logger.Logger, st *storage.Storage) *Server {
	s := &Server{
		log: log,
		db:  st.Temp,
	}
	fmt.Println("Server load")
	return s
}

func (s *Server) RunServer() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// Обработчик для принятия сообщений от DiscordService
	router.OPTIONS("/compendium/applink/identities", s.CheckIdentityOptions)
	router.GET("/compendium/applink/identities", s.CheckIdentityHandler)
	router.OPTIONS("/compendium/applink/connect", s.CheckConnectOptions)
	router.POST("/compendium/applink/connect", s.CheckConnectHandler)
	router.OPTIONS("/compendium/cmd/syncTech/get", s.CheckSyncTechOptions)
	router.POST("/compendium/cmd/syncTech/get", s.CheckSyncTechHandler)
	router.POST("/cmd/syncTech", s.SyncHandler)
	router.OPTIONS("/compendium/cmd/corpdata", s.Check)            //?roleId=
	router.GET("/compendium/cmd/corpdata", s.CheckCorpDataHandler) //?roleId=

	err := router.Run(":80")
	if err != nil {
		s.log.ErrorErr(err)
		os.Exit(1)
	}
}
