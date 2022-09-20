package services

import (
	//	"errors"

	"postoffice/app/core"
	"postoffice/app/models"
	"postoffice/app/pkg"
	"postoffice/app/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
	//	"gorm.io/gorm"
)

type appServiceLayer struct {
	repository repository.Repo
	config     *core.Config
}

func newAppServiceLayer(r repository.Repo, c *core.Config) *appServiceLayer {
	return &appServiceLayer{
		repository: r,
		config:     c,
	}
}

func (a *appServiceLayer) CreateApp(req core.CreateAppRequest) core.Response {
	app := models.App{
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
		ApiKey:      pkg.GenerateApiKey(),
	}
	if err := a.repository.Apps.Create(&app); err != nil {
		return core.Error(err, nil)
	}

	return core.Success(&map[string]interface{}{
		"app": app,
	}, core.String("app created successfully"))
}

func (a *appServiceLayer) FetchApps() core.Response {

	var apps []models.App

	err := a.repository.Apps.Fetch(&apps)
	if err != nil {
		return core.Error(err, nil)
	}
	if len(apps) < 1 {
		return core.NoContentFound(err, core.String("No apps found"))
	}

	return core.Success(&map[string]interface{}{
		"apps": apps,
	}, core.String("apps found successfully"))
}

func (a *appServiceLayer) GetApp(id string) core.Response {
	app := models.App{}
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return core.Error(err, nil)
	}

	if err := a.repository.Apps.Get(&app, objectId); err != nil {
		return core.BadRequest(err, nil)
	}

	var modules []models.Module

	a.repository.Modules.FetchByAppID(&modules, objectId)
	app.Modules = modules
	return core.Success(&map[string]interface{}{
		"app": app,
	}, core.String("app found successfully"))
}

func (a *appServiceLayer) UpdateApp(req core.UpdateAppRequest) core.Response {
	app := models.App{}
	objectId, err := primitive.ObjectIDFromHex(req.Id)

	if err != nil {
		return core.Error(err, nil)
	}

	if err := a.repository.Apps.Get(&app, objectId); err != nil {
		return core.BadRequest(err, core.String("App does not exist"))
	}

	app.Name = req.Name
	app.Status = req.Status
	app.Description = req.Description

	if err := a.repository.Apps.Update(&app); err != nil {
		return core.Error(err, nil)
	}

	return core.Success(&map[string]interface{}{
		"app": app,
	}, core.String("app updated successfully"))
}
