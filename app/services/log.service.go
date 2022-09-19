package services

import (
	//	"errors"
	"postoffice/app/core"
	"postoffice/app/models"
	"postoffice/app/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
	//	"gorm.io/gorm"
)

type logServiceLayer struct {
	repository repository.Repo
	config     *core.Config
}

func newLogServiceLayer(r repository.Repo, c *core.Config) *appServiceLayer {
	return &appServiceLayer{
		repository: r,
		config:     c,
	}
}

func (a *logServiceLayer) CreateApp(req core.CreateLogRequest) core.Response {
	domainId, err := primitive.ObjectIDFromHex(req.DomainId)

	if err != nil {
		return core.Error(err, nil)
	}

	log := models.Log{
		Data:     req.Data,
		DomainId: domainId,
		Action:   req.Action,
		Creator:  req.Creator,
	}
	if err := a.repository.Logs.Create(&log); err != nil {
		return core.Error(err, nil)
	}

	return core.Success(&map[string]interface{}{
		"log": log,
	}, core.String("log created successfully"))
}
