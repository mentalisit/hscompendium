package mongo

import (
	"compendium/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

func (d *DB) ReadSyncData(token string) models.SyncData {
	var data models.SyncData
	//var data models.SyncData
	collection := d.s.Database("Compendium").Collection("SyncData")
	err := collection.FindOne(context.Background(), bson.M{"token": token}).Decode(&data)
	if err != nil {
		d.log.ErrorErr(err)
		return models.SyncData{}
	}
	return data
}
func (d *DB) InsertSyncData(s models.SyncData) {

	d.s.Database("Compendium").CreateCollection(context.Background(), "SyncData")

	collection := d.s.Database("Compendium").Collection("SyncData")
	ins, err := collection.InsertOne(context.Background(), s)
	if err != nil {
		d.log.ErrorErr(err)
	}
	fmt.Println(ins.InsertedID)
}
func (d *DB) UpdateSyncData(s models.SyncData) {
	collection := d.s.Database("Compendium").Collection("SyncData")
	filter := bson.M{"token": s.Token}
	_, err := collection.ReplaceOne(context.Background(), filter, s)
	if err != nil {
		d.log.ErrorErr(err)
	}
}
