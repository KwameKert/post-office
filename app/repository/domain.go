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

type domainLayer struct {
	collection *mongo.Collection
}

func newDomainRepoLayer(db *mongo.Database) *domainLayer {
	return &domainLayer{
		collection: db.Collection("Domains"),
	}
}

func (dl *domainLayer) Create(domain *models.Domain) error {
	domain.CreatedAt = time.Now()
	_, err := dl.collection.InsertOne(cxt.TODO(), &domain)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (dl *domainLayer) Update(domain *models.Domain) error {
	filter := bson.D{{"_id", domain.Id}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{"name", domain.Name},
		{"description", domain.Description},
		{"status", domain.Status},
		{"module", domain.Module},
		{"updated_at", time.Now()},
	}}}

	_, err := dl.collection.UpdateOne(cxt.TODO(), filter, update)

	if err != nil {
		return err
	}
	return nil
}

func (dl *domainLayer) Get(domain *bson.M, id primitive.ObjectID) error {

	//query := bson.M{"_id": id}
	lookupStage := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "Modules"},
			{Key: "localField", Value: "module_id"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "module"},
		}}}
	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: id}}}}

	showInfoCursor, err := dl.collection.Aggregate(cxt.TODO(), mongo.Pipeline{matchStage, lookupStage})
	if err != nil {
		panic(err)
	}

	for showInfoCursor.Next(cxt.TODO()) {
		if err := showInfoCursor.Decode(&domain); err != nil {
			log.Fatal(err)
		}
	}
	if err := showInfoCursor.Err(); err != nil {
		log.Fatal(err)
	}
	return nil
}

func (dl *domainLayer) FetchByModuleID(domains *[]models.Domain, id primitive.ObjectID) error {

	cursor, err := dl.collection.Find(context.TODO(), bson.M{"module_id": id})
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(context.TODO(), domains); err != nil {
		log.Fatal(err)
	}
	return nil
}

func (dl *domainLayer) Fetch(domains *[]models.Domain) error {
	// lookupStage := bson.D{
	// 	{Key: "$lookup", Value: bson.D{
	// 		{Key: "from", Value: "Modules"},
	// 		{Key: "localField", Value: "module_id"},
	// 		{Key: "foreignField", Value: "_id"},
	// 		{Key: "as", Value: "modules"},
	// 	}}}

	//cursor, err := dl.collection.Aggregate(cxt.TODO(), mongo.Pipeline{lookupStage})

	cursor, err := dl.collection.Find(cxt.TODO(), bson.D{})

	if err != nil {
		panic(err)
	}

	if err = cursor.All(context.TODO(), domains); err != nil {
		log.Fatal(err)
	}
	return nil
}
