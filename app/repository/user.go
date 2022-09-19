package repository

import (
	//	"errors"
	cxt "context"
	"postoffice/app/models"
	"time"

	// log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userLayer struct {
	collection *mongo.Collection
}

func newUserRepoLayer(db *mongo.Database) *userLayer {
	return &userLayer{
		collection: db.Collection("Users"),
	}
}

func (ul *userLayer) Create(user *models.User) error {
	user.CreatedAt = time.Now()
	_, err := ul.collection.InsertOne(cxt.TODO(), &user)
	if err != nil {
		return err
	}
	return nil
}

func (ul *userLayer) Update(user *models.User) error {
	filter := bson.D{{"_id", user.Id}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{"username", user.Username},
		{"email", user.Email},
		{"description", user.Description},
		{"updated_at", time.Now()},
	}}}

	_, err := ul.collection.UpdateOne(cxt.TODO(), filter, update)

	if err != nil {
		return err
	}
	return nil
}

func (ul *userLayer) Fetch() (error, []models.User) {

	cursor, err := ul.collection.Find(cxt.TODO(), bson.D{})
	if err != nil {
		return err, nil
	}

	var results []models.User

	if err = cursor.All(cxt.TODO(), &results); err != nil {
		panic(err)
	}
	return nil, results
}

func (ul *userLayer) Get(user *models.User, id primitive.ObjectID) error {

	query := bson.M{"_id": id}
	if err := ul.collection.FindOne(cxt.TODO(), query).Decode(&user); err != nil {
		return err
	}
	return nil
}

func (ul *userLayer) GetByUsername(user *models.User, username string) error {
	query := bson.M{"username": username}
	if err := ul.collection.FindOne(cxt.TODO(), query).Decode(&user); err != nil {
		return err
	}
	return nil
}
