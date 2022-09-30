package repository

import (
	//	"errors"
	cxt "context"
	"postoffice/app/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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
	log.CreatedAt = time.Now()
	log.UpdatedAt = time.Now()
	_, err := al.collection.InsertOne(cxt.TODO(), &log)
	if err != nil {
		return err
	}
	return nil
}

func (dl *logLayer) Search(logs *[]bson.M) error {

	//query := bson.M{"_id": id}
	pipelines := mongo.Pipeline{
		{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "Domains"},
				{Key: "localField", Value: "domain"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "domain"},
			}},
		},
		{{Key: "$unwind", Value: "$domain"}},
		{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "Modules"},
				{Key: "localField", Value: "domain.module"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "module"},
			}},
		},
		{{Key: "$unwind", Value: "$module"}},
	}

	// lookupStage := bson.D{
	// 	{Key: "$lookup", Value: bson.D{
	// 		{Key: "from", Value: "Domains"},
	// 		{Key: "localField", Value: "domain"},
	// 		{Key: "foreignField", Value: "_id"},
	// 		{Key: "as", Value: "domain"},
	// 	}}}

	// lookupStage2 := bson.D{{Key: "$lookup", Value: bson.D{
	// 	{Key: "from", Value: "Module"},
	// 	{Key: "localField", Value: "module"},
	// 	{Key: "foreignField", Value: "_id"},
	// 	{Key: "as", Value: "module"},
	// }}}

	showInfoCursor, err := dl.collection.Aggregate(cxt.TODO(), pipelines)
	if err != nil {
		panic(err)
	}

	if err := showInfoCursor.All(cxt.TODO(), logs); err != nil {
		panic(err)
	}

	return nil
}
