package server

import (
	"compendium/models"
	"compendium/storage"
	"fmt"
	"github.com/mentalisit/logger"
	"log"
	"net/http"
)

type Server struct {
	log *logger.Logger
	db  db
}
type db interface {
	ReadIdentity() []models.Identity
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
	// Обработчик запросов для проверки идентификации
	http.HandleFunc("/compendium/applink/identities", s.CheckIdentityHandler)
	http.HandleFunc("/cmd/syncTech", s.SyncHandler)

	// Запуск HTTP-сервера на порту 8080
	log.Fatal(http.ListenAndServe(":80", nil))
}
