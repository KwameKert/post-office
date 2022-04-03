package services

import (
	"postoffice/app/core"
	"postoffice/app/repository"
)

type Services struct {
	UserService *userServiceLayer
	AppService  *appServiceLayer
}

func NewService(r repository.Repo, c *core.Config) Services {
	return Services{
		UserService: newUserServiceLayer(r, c),
		AppService:  newAppServiceLayer(r, c),
	}
}
