package mongo

import (
	"compendium/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

func (d *DB) ReadIdentity() []models.Identity {
	collection := d.s.Database("Compendium").Collection("Identity")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		d.log.ErrorErr(err)
		return nil
	}
	var m []models.Identity
	err = cursor.All(context.Background(), &m)
	if err != nil {
		d.log.ErrorErr(err)
		return nil
	}
	return m
}
func (d *DB) ReadIdentityByToken(token string) models.Identity {
	var data models.Identity
	collection := d.s.Database("Compendium").Collection("Identity")
	err := collection.FindOne(context.Background(), bson.M{"token": token}).Decode(&data)
	if err != nil {
		d.log.ErrorErr(err)
		return models.Identity{}
	}
	return data
}

func (d *DB) InsertIdentity(c models.Identity) {

	d.s.Database("Compendium").CreateCollection(context.Background(), "Identity")

	collection := d.s.Database("Compendium").Collection("Identity")
	ins, err := collection.InsertOne(context.Background(), c)
	if err != nil {
		d.log.ErrorErr(err)
	}
	fmt.Println(ins.InsertedID)
}

func (d *DB) UpdateIdentity(c models.Identity) {
	collection := d.s.Database("Compendium").Collection("Identity")
	filter := bson.M{"token": c.Token}
	_, err := collection.ReplaceOne(context.Background(), filter, c)
	if err != nil {
		d.log.ErrorErr(err)
	}
}
