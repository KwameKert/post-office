package services

import (
	"postoffice/app/core"
	"postoffice/app/repository"
)

type Services struct {
	UserService   *userServiceLayer
	AppService    *appServiceLayer
	DomainService *domainServiceLayer
	ModuleService *moduleServiceLayer
}

func NewService(r repository.Repo, c *core.Config) Services {
	return Services{
		UserService:   newUserServiceLayer(r, c),
		AppService:    newAppServiceLayer(r, c),
		DomainService: newDomainServiceLayer(r, c),
		ModuleService: newModuleServiceLayer(r, c),
	}
}
