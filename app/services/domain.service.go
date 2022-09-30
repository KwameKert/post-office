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

	//add app to domain

	domain := models.Domain{
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
		Module:      moduleId,
	}
	if err := a.repository.Domains.Create(&domain); err != nil {
		return core.Error(err, nil)
	}

	return core.Success(&map[string]interface{}{
		"domain": domain,
	}, core.String("domain created successfully"))
}

func (d *domainServiceLayer) GetDomain(id string) core.Response {
	domain := bson.M{}
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return core.Error(err, nil)
	}

	if err := d.repository.Domains.Get(&domain, objectId); err != nil {
		return core.BadRequest(err, nil)
	}

	//domainResponse := formatDomainResponse(&domain)

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
	domainFound := bson.M{}
	if err := d.repository.Domains.Get(&domainFound, domainId); err != nil {
		return core.BadRequest(err, core.String("Domain does not exist"))
	}

	//add app to domain

	moduleId, err := primitive.ObjectIDFromHex(req.ModuleId)

	if err != nil {
		return core.Error(err, nil)
	}

	domain.Id = domainId
	domain.Name = req.Name
	domain.Status = req.Status
	domain.Description = req.Description
	domain.Module = moduleId

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

func formatDomainResponse(domainBson *bson.M) core.DomainResponse {
	domainDecoded := core.TempDomain{}
	bsonBytes, _ := bson.Marshal(domainBson)
	bson.Unmarshal(bsonBytes, &domainDecoded)
	domain := core.DomainResponse{}
	domain.CreatedAt = domainDecoded.CreatedAt
	domain.UpdatedAt = domainDecoded.UpdatedAt
	domain.Name = domainDecoded.Name
	domain.Description = domainDecoded.Description
	domain.Status = domainDecoded.Status
	domain.Module = domainDecoded.Module[0]

	return domain
}
