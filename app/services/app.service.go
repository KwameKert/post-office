package services

import (
	//	"errors"
	"postoffice/app/core"
	"postoffice/app/models"
	"postoffice/app/repository"

	"go.mongodb.org/mongo-driver/bson"
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
	}
	println("Im here")
	if err := a.repository.Apps.Create(&app); err != nil {
		return core.Error(err, nil)
	}

	return core.Success(&map[string]interface{}{
		"app": app,
	}, core.String("app created successfully"))
}

func (a *appServiceLayer) FetchApps() core.Response {

	var apps []bson.M

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
