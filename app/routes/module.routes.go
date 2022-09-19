package routes

import (
	"net/http"
	"postoffice/app/core"
	"postoffice/app/middlewares"
	"postoffice/app/services"

	"github.com/gin-gonic/gin"
)

func RegisterModuleRoutes(e *gin.Engine, s services.Services) {
	serve = s
	e.POST(apiUrl+"/modules", middlewares.AuthorizeClientRequest(), createModule)
	e.GET(apiUrl+"/modules", middlewares.AuthorizeClientRequest(), fetchModules)
	e.GET(apiUrl+"/modules/:id", middlewares.AuthorizeClientRequest(), getModules)
	e.PUT(apiUrl+"/modules", middlewares.AuthorizeClientRequest(), updateModule)
}

func createModule(c *gin.Context) {
	var req core.CreateModuleRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	response := serve.ModuleService.CreateModule(req)

	if response.Error {
		c.JSON(response.Code, gin.H{
			"message": response.Meta.Message,
		})
		return
	}

	c.JSON(response.Code, response.Meta)
}

func updateModule(c *gin.Context) {
	var req core.UpdateModuleRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	response := serve.ModuleService.UpdateModule(req)

	if response.Error {
		c.JSON(response.Code, gin.H{
			"message": response.Meta.Message,
		})
		return
	}

	c.JSON(response.Code, response.Meta)
}

func fetchModules(c *gin.Context) {

	response := serve.ModuleService.FetchModules()
	if response.Error {
		c.JSON(response.Code, gin.H{
			"message": response.Meta.Message,
		})
		return
	}

	c.JSON(response.Code, response.Meta)
}

func getModules(c *gin.Context) {

	idStr := c.Param("id")
	response := serve.ModuleService.GetModule(idStr)
	if response.Error {
		c.JSON(response.Code, gin.H{
			"message": response.Meta.Message,
		})
		return
	}

	c.JSON(response.Code, response.Meta)
}
