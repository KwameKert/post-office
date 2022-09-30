package services

import (
	//	"errors"
	"postoffice/app/core"
	"postoffice/app/models"
	"postoffice/app/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//	"gorm.io/gorm"
)

type logServiceLayer struct {
	repository repository.Repo
	config     *core.Config
}

func newLogServiceLayer(r repository.Repo, c *core.Config) *logServiceLayer {
	return &logServiceLayer{
		repository: r,
		config:     c,
	}
}

func (a *logServiceLayer) CreateLog(req core.CreateLogRequest) core.Response {
	domainId, err := primitive.ObjectIDFromHex(req.DomainId)

	if err != nil {
		return core.Error(err, nil)
	}

	//add module, app , domain
	domain := bson.M{}
	if err := a.repository.Domains.Get(&domain, domainId); err != nil {
		return core.BadRequest(err, core.String("No domain found"))
	}

	log := models.Log{
		Data:   req.Data,
		Domain: domainId,
		//ModuleId: ,
		Action:  req.Action,
		Creator: req.UserId,
	}
	if err := a.repository.Logs.Create(&log); err != nil {
		return core.Error(err, nil)
	}

	return core.Success(&map[string]interface{}{
		"log": log,
	}, core.String("log created successfully"))
}

func (a *logServiceLayer) SearchLog() core.Response {
	logs := []bson.M{}
	if err := a.repository.Logs.Search(&logs); err != nil {
		return core.BadRequest(err, core.String("No logs"))
	}

	return core.Success(&map[string]interface{}{
		"logs": logs,
	}, core.String("logs found successfully"))
}
