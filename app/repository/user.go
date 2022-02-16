package repository

import (
	//	"errors"
	cxt "context"
	"postoffice/app/models"

	// log "github.com/sirupsen/logrus"
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
	_, err := ul.collection.InsertOne(cxt.TODO(), &user)
	if err != nil {
		return err
	}
	return nil
}

func (ul *userLayer) Fetch(user *[]models.User) error {

	// if err := ul.db.Find(&user).Error; err != nil {
	// 	log.Error("error -->", err)
	// 	return err
	// }
	return nil
}

func (ul *userLayer) Get(user *models.User, id int) error {

	// if err := ul.db.Preload("Tasks").Find(&user, id).First(&user).Error; err != nil {

	// 	return err
	// }
	return nil
}
