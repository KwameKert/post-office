package repository

import (
	//	"errors"
	"context"
	cxt "context"
	"postoffice/app/models"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type appLayer struct {
	collection *mongo.Collection
}

func newAppRepoLayer(db *mongo.Database) *appLayer {
	return &appLayer{
		collection: db.Collection("Apps"),
	}
}

func (al *appLayer) Create(app *models.App) error {
	_, err := al.collection.InsertOne(cxt.TODO(), &app)
	if err != nil {
		return err
	}
	return nil
}

func (al *appLayer) Fetch(apps *[]bson.M) error {

	cursor, err := al.collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(context.TODO(), apps); err != nil {
		log.Fatal(err)
	}
	return nil
}
