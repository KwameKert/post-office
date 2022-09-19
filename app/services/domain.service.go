package services

import (
	//	"errors"
	"postoffice/app/core"
	"postoffice/app/models"
	"postoffice/app/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
	//	"gorm.io/gorm"
)

type domainServiceLayer struct {
	repository repository.Repo
	config     *core.Config
}

func newDomainServiceLayer(r repository.Repo, c *core.Config) *domainServiceLayer {
	return &domainServiceLayer{
		repository: r,
		config:     c,
	}
}

func (a *domainServiceLayer) CreateDomain(req core.CreateDomainRequest) core.Response {
	moduleId, err := primitive.ObjectIDFromHex(req.ModuleId)

	if err != nil {
		return core.Error(err, nil)
	}

	domain := models.Domain{
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
		ModuleId:    moduleId,
	}
	if err := a.repository.Domains.Create(&domain); err != nil {
		return core.Error(err, nil)
	}

	return core.Success(&map[string]interface{}{
		"domain": domain,
	}, core.String("domain created successfully"))
}

func (d *domainServiceLayer) GetDomain(id string) core.Response {
	domain := models.Domain{}
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return core.Error(err, nil)
	}

	if err := d.repository.Domains.Get(&domain, objectId); err != nil {
		return core.BadRequest(err, nil)
	}

	return core.Success(&map[string]interface{}{
		"domain": domain,
	}, core.String("domain found successfully"))
}

func (d *domainServiceLayer) UpdateDomain(req core.UpdateDomainRequest) core.Response {
	domain := models.Domain{}
	domainId, err := primitive.ObjectIDFromHex(req.Id)

	if err != nil {
		return core.Error(err, nil)
	}

	if err := d.repository.Domains.Get(&domain, domainId); err != nil {
		return core.BadRequest(err, core.String("Domain does not exist"))
	}

	moduleId, err := primitive.ObjectIDFromHex(req.ModuleId)

	domain.Name = req.Name
	domain.Status = req.Status
	domain.Description = req.Description
	domain.ModuleId = moduleId

	if err := d.repository.Domains.Update(&domain); err != nil {
		return core.Error(err, nil)
	}

	return core.Success(&map[string]interface{}{
		"domain": domain,
	}, core.String("domain updated successfully"))
}

func (a *domainServiceLayer) FetchDomains() core.Response {

	var domains []models.Domain

	err := a.repository.Domains.Fetch(&domains)
	if err != nil {
		return core.Error(err, nil)
	}
	if len(domains) < 1 {
		return core.NoContentFound(err, core.String("No domains found"))
	}

	return core.Success(&map[string]interface{}{
		"domains": domains,
	}, core.String("domains found successfully"))
}
