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
		{"module_id", domain.ModuleId},
		{"updated_at", time.Now()},
	}}}

	_, err := dl.collection.UpdateOne(cxt.TODO(), filter, update)

	if err != nil {
		return err
	}
	return nil
}

func (dl *domainLayer) Get(domain *models.Domain, id primitive.ObjectID) error {

	query := bson.M{"_id": id}
	if err := dl.collection.FindOne(cxt.TODO(), query).Decode(&domain); err != nil {
		return err
	}
	return nil
}

func (dl *domainLayer) Fetch(domains *[]models.Domain) error {

	cursor, err := dl.collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(context.TODO(), domains); err != nil {
		log.Fatal(err)
	}
	return nil
}
