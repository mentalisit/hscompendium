package mongo

import (
	"compendium/config"
	"context"
	"github.com/mentalisit/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	s   *mongo.Client
	log *logger.Logger
}

func InitMongoDB(log *logger.Logger) *DB {
	uri := config.Instance.Mongo
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic("mongo")
	}
	d := &DB{
		s:   client,
		log: log,
	}
	return d
}
