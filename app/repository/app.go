package repository

import (
	//	"errors"
	"context"
	cxt "context"
	"postoffice/app/models"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	app.CreatedAt = time.Now()
	_, err := al.collection.InsertOne(cxt.TODO(), &app)
	if err != nil {
		return err
	}
	return nil
}

func (al *appLayer) Fetch(apps *[]models.App) error {

	cursor, err := al.collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(context.TODO(), apps); err != nil {
		log.Fatal(err)
	}
	return nil
}

func (al *appLayer) Update(app *models.App) error {
	filter := bson.D{{"_id", app.Id}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{"name", app.Name},
		{"description", app.Description},
		{"status", app.Status},
		{"updated_at", time.Now()},
	}}}

	_, err := al.collection.UpdateOne(cxt.TODO(), filter, update)

	if err != nil {
		return err
	}
	return nil
}

func (al *appLayer) Get(app *models.App, id primitive.ObjectID) error {

	query := bson.M{"_id": id}
	if err := al.collection.FindOne(cxt.TODO(), query).Decode(&app); err != nil {
		return err
	}
	return nil
}
