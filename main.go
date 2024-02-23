package main

import (
	"compendium/Compendium"
	"compendium/client/ds"
	"compendium/client/server"
	"compendium/client/tg"
	"compendium/config"
	"compendium/storage"
	"fmt"
	"github.com/mentalisit/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.InitConfig()
	log := logger.LoggerZapDEV()
	st := storage.NewStorage(log, cfg)

	d := ds.NewDiscord(log, st, cfg)
	t := tg.NewTelegram(log, cfg, st)

	Compendium.NewCompendium(d, t, log)
	s := server.NewServer(log, st)
	s.RunServer()

	fmt.Println("load ok ")
	//ожидаем сигнала завершения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}
