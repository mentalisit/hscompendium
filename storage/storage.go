package storage

import (
	"compendium/config"
	"compendium/storage/mongo"
	"github.com/mentalisit/logger"
	"go.uber.org/zap"
)

type Storage struct {
	log   *zap.Logger
	debug bool
	Temp  *mongo.DB
}

func NewStorage(log *logger.Logger, cfg *config.ConfigBot) *Storage {

	mongoDB := mongo.InitMongoDB(log)

	s := &Storage{

		Temp: mongoDB,
	}

	//go s.loadDbArray()

	return s
}
