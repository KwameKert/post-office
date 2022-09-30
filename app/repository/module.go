package repository

import (
	//	"errors"
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
	module.UpdatedAt = time.Now()
	_, err := ml.collection.InsertOne(cxt.TODO(), &module)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (ml *moduleLayer) Fetch(modules *[]models.Module) error {

	// lookupStage := bson.D{
	// 	{Key: "$lookup", Value: bson.D{
	// 		{Key: "from", Value: "Apps"},
	// 		{Key: "localField", Value: "app"},
	// 		{Key: "foreignField", Value: "_id"},
	// 		{Key: "as", Value: "app"},
	// 	}}}
	// cursor, err := ml.collection.Aggregate(cxt.TODO(), mongo.Pipeline{lookupStage})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if err = cursor.All(cxt.TODO(), modules); err != nil {
	// 	log.Fatal(err)
	// }

	cursor, err := ml.collection.Find(cxt.TODO(), bson.D{})

	if err != nil {
		panic(err)
	}

	if err = cursor.All(cxt.TODO(), modules); err != nil {
		log.Fatal(err)
	}

	return nil
}

func (ml *moduleLayer) FetchByAppID(modules *[]models.Module, id primitive.ObjectID) error {

	cursor, err := ml.collection.Find(cxt.TODO(), bson.M{"app_id": id})
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(cxt.TODO(), modules); err != nil {
		log.Fatal(err)
	}
	return nil
}

func (ml *moduleLayer) Update(module *models.Module) error {
	filter := bson.D{{Key: "_id", Value: module.Id}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "name", Value: module.Name},
		{Key: "description", Value: module.Description},
		{Key: "status", Value: module.Status},
		{Key: "app_id", Value: module.App},
		{Key: "updated_at", Value: time.Now()},
	}}}

	_, err := ml.collection.UpdateOne(cxt.TODO(), filter, update)

	if err != nil {
		return err
	}
	return nil
}

func (ml *moduleLayer) Get(module *bson.M, id primitive.ObjectID) error {

	lookupStage := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "Apps"},
			{Key: "localField", Value: "app"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "app"},
		}}}
	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: id}}}}

	showInfoCursor, err := ml.collection.Aggregate(cxt.TODO(), mongo.Pipeline{matchStage, lookupStage})
	if err != nil {
		panic(err)
	}

	for showInfoCursor.Next(cxt.TODO()) {
		if err := showInfoCursor.Decode(&module); err != nil {
			log.Fatal(err)
		}
	}
	if err := showInfoCursor.Err(); err != nil {
		log.Fatal(err)
	}
	return nil
}
