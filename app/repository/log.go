package repository

import (
	//	"errors"
	//	"context"
	cxt "context"

	log "github.com/sirupsen/logrus"

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

func (al *logLayer) AddIndex(data *bson.D) {
	model := mongo.IndexModel{Keys: bson.D{{"data", "text"}}}
	name, err := al.collection.Indexes().CreateOne(cxt.TODO(), model)
	if err != nil {
		panic(err)
	}
	log.Info("name is here {}", name)
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

func (dl *logLayer) Search(logs *[]bson.M, query bson.D, limit int, skip int) error {

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
		query,
		{{
			Key: "$setWindowFields", Value: bson.D{{
				Key: "output", Value: bson.D{{
					Key: "totalCount", Value: bson.D{{"$count", bson.D{}}},
				}},
			}},
		}},
		{{Key: "$skip", Value: skip}},
		{{Key: "$limit", Value: limit}},
	}

	showInfoCursor, err := dl.collection.Aggregate(cxt.TODO(), pipelines)
	if err != nil {
		log.Panic(err)
	}

	if err := showInfoCursor.All(cxt.TODO(), logs); err != nil {
		log.Panic(err)
	}

	return nil
}
