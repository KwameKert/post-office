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

func (a *logServiceLayer) SearchLog(data core.SearchRequest) core.Response {
	logs := []bson.M{}

	query := a.searchQueryBuilder(&data)
	if err := a.repository.Logs.Search(&logs, query); err != nil {
		return core.BadRequest(err, core.String("No logs"))
	}

	return core.Success(&map[string]interface{}{
		"logs": logs,
	}, core.String("logs found successfully"))
}

func (a *logServiceLayer) searchQueryBuilder(data *core.SearchRequest) bson.D {

	var queryArr []bson.D

	queryArr = append(queryArr, bson.D{{"action", data.Action}})

	if data.DomainId != "" {
		domainId, _ := primitive.ObjectIDFromHex(data.DomainId)
		queryArr = append(queryArr, bson.D{{Key: "domain._id", Value: domainId}})
	}

	if data.ModuleId != "" {
		moduleId, _ := primitive.ObjectIDFromHex(data.ModuleId)
		queryArr = append(queryArr, bson.D{{Key: "module._id", Value: moduleId}})

	}

	if data.UserId != "" {
		queryArr = append(queryArr, bson.D{{Key: "user_id", Value: data.UserId}})
	}
	// if data.Text != "" {
	// 	indexField := bson.D{{Key: "data", Value: "text"}}
	// 	a.repository.Logs.AddIndex(&indexField)

	// 	queryData = append(queryData, bson.E{Key: "$text",
	// 		Value: bson.D{{Key: "$search", Value: data.Text}}})
	// }
	// array := []bson.D{}
	// array = append(array, bson.D{{"user_id", "632309ac267f3818b3ad5071"}})
	// array = append(array, bson.D{{"action", "INSERT"}})

	matchStage := bson.D{{
		"$match", bson.D{{"$and", queryArr}}},
	}
	return matchStage

}
