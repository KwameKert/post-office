package routes

import (
	"postoffice/app/core"
	"postoffice/app/services"

	"github.com/gin-gonic/gin"
)

type router struct {
	engine   *gin.Engine
	config   *core.Config
	services services.Services
}

const apiUrl string = "/api/v1"

var serve services.Services

func NewRouter(engine *gin.Engine, config *core.Config, services services.Services) *router {
	return &router{
		engine:   engine,
		config:   config,
		services: services,
	}
}

func (r *router) RegisterRoutes() {
	RegisterUserRoutes(r.engine, r.services)
	RegisterAppRoutes(r.engine, r.services)
	RegisterModuleRoutes(r.engine, r.services)
	RegisterDomainRoutes(r.engine, r.services)
}
