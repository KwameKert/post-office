package repository

import (
	//	"errors"
	"postoffice/app/models"

	//log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type issueLayer struct {
	collection *mongo.Collection
}

func newIssueRepoLayer(db *mongo.Database) *issueLayer {
	return &issueLayer{
		collection: db.Collection("Users"),
	}
}

func (ul *issueLayer) Create(issue *models.Issue) error {

	// if err := ul.db.(issue).Error; err != nil {
	// 	return err
	// }
	return nil
}

func (ul *issueLayer) Fetch(issue *[]models.Issue) error {

	// if err := ul.db.Find(&issue).Error; err != nil {
	// 	log.Error("error -->", err)
	// 	return err
	// }
	return nil
}
