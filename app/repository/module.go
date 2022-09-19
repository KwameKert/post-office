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

type moduleLayer struct {
	collection *mongo.Collection
}

func newModuleRepoLayer(db *mongo.Database) *moduleLayer {
	return &moduleLayer{
		collection: db.Collection("Modules"),
	}
}

func (ml *moduleLayer) Create(module *models.Module) error {
	module.CreatedAt = time.Now()
	_, err := ml.collection.InsertOne(cxt.TODO(), &module)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (ml *moduleLayer) Fetch(modules *[]models.Module) error {

	cursor, err := ml.collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(context.TODO(), modules); err != nil {
		log.Fatal(err)
	}
	return nil
}

func (ml *moduleLayer) FetchByAppID(modules *[]models.Module, id primitive.ObjectID) error {

	cursor, err := ml.collection.Find(context.TODO(), bson.M{"app_id": id})
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(context.TODO(), modules); err != nil {
		log.Fatal(err)
	}
	return nil
}

func (ml *moduleLayer) Update(module *models.Module) error {
	filter := bson.D{{"_id", module.Id}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{"name", module.Name},
		{"description", module.Description},
		{"status", module.Status},
		{"app_id", module.AppID},
		{"updated_at", time.Now()},
	}}}

	_, err := ml.collection.UpdateOne(cxt.TODO(), filter, update)

	if err != nil {
		return err
	}
	return nil
}

func (ml *moduleLayer) Get(module *models.Module, id primitive.ObjectID) error {

	query := bson.M{"_id": id}
	if err := ml.collection.FindOne(cxt.TODO(), query).Decode(&module); err != nil {
		return err
	}
	return nil
}
