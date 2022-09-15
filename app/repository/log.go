package repository

import (
	//	"errors"
	cxt "context"
	"postoffice/app/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type logLayer struct {
	collection *mongo.Collection
}

func newLogRepoLayer(db *mongo.Database) *logLayer {
	return &logLayer{
		collection: db.Collection("Logs"),
	}
}

func (al *logLayer) Create(log *models.Log) error {
	_, err := al.collection.InsertOne(cxt.TODO(), &log)
	if err != nil {
		return err
	}
	return nil
}
