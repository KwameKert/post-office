package services

import (
	//	"errors"
	"fmt"
	"postoffice/app/core"
	"postoffice/app/models"
	"postoffice/app/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//	"gorm.io/gorm"
)

type moduleServiceLayer struct {
	repository repository.Repo
	config     *core.Config
}

func newModuleServiceLayer(r repository.Repo, c *core.Config) *moduleServiceLayer {
	return &moduleServiceLayer{
		repository: r,
		config:     c,
	}
}

func (m *moduleServiceLayer) CreateModule(req core.CreateModuleRequest) core.Response {
	appId, err := primitive.ObjectIDFromHex(req.AppId)

	if err != nil {
		return core.Error(err, nil)
	}

	//fetch app by Id
	var app *models.App
	if err := m.repository.Apps.Get(app, appId); err != nil {
		return core.Error(err, nil)
	}

	module := models.Module{
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
		App:         appId,
	}
	if err := m.repository.Modules.Create(&module); err != nil {
		return core.Error(err, nil)
	}

	return core.Success(&map[string]interface{}{
		"module": module,
	}, core.String("module created successfully"))
}

func (m *moduleServiceLayer) GetModule(id string) core.Response {
	module := bson.M{}
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return core.Error(err, nil)
	}

	if err := m.repository.Modules.Get(&module, objectId); err != nil {
		return core.BadRequest(err, nil)
	}

	if module == nil || len(module) == 0 {
		return core.BadRequest(err, core.String("No module found"))
	}
	fmt.Println(module)

	//moduleResponse := formatModuleResponse(&module)

	return core.Success(&map[string]interface{}{
		"module": module,
	}, core.String("domain found successfully"))
}

func (m *moduleServiceLayer) UpdateModule(req core.UpdateModuleRequest) core.Response {
	module := models.Module{}
	moduleId, err := primitive.ObjectIDFromHex(req.Id)

	if err != nil {
		return core.Error(err, nil)
	}

	moduleFound := bson.M{}
	if err := m.repository.Modules.Get(&moduleFound, moduleId); err != nil {
		return core.BadRequest(err, core.String("Module does not exist"))
	}

	appId, err := primitive.ObjectIDFromHex(req.Id)

	if err != nil {
		return core.Error(err, nil)
	}

	var app *models.App
	if err := m.repository.Apps.Get(app, appId); err != nil {
		return core.Error(err, nil)
	}
	module.Id = moduleId
	module.Name = req.Name
	module.Status = req.Status
	module.Description = req.Description
	module.App = appId

	if err := m.repository.Modules.Update(&module); err != nil {
		return core.Error(err, nil)
	}

	return core.Success(&map[string]interface{}{
		"module": module,
	}, core.String("module updated successfully"))
}

func (m *moduleServiceLayer) FetchModules() core.Response {

	var modules []models.Module

	err := m.repository.Modules.Fetch(&modules)
	if err != nil {
		return core.Error(err, nil)
	}
	if len(modules) < 1 {
		return core.NoContentFound(err, core.String("No modules found"))
	}

	return core.Success(&map[string]interface{}{
		"modules": modules,
	}, core.String("modules found successfully"))
}

func formatModuleResponse(moduleBson *bson.M) core.ModuleResponse {
	moduleDecoded := core.TempModule{}
	bsonBytes, _ := bson.Marshal(moduleBson)
	bson.Unmarshal(bsonBytes, &moduleDecoded)
	module := core.ModuleResponse{}
	module.CreatedAt = moduleDecoded.CreatedAt
	module.UpdatedAt = moduleDecoded.UpdatedAt
	module.Name = moduleDecoded.Name
	module.Description = moduleDecoded.Description
	module.Status = moduleDecoded.Status
	module.App = moduleDecoded.App[0]

	return module
}
