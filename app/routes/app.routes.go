package routes

import (
	"net/http"
	"postoffice/app/core"
	"postoffice/app/middlewares"
	"postoffice/app/services"

	"github.com/gin-gonic/gin"
)

func RegisterAppRoutes(e *gin.Engine, s services.Services) {
	serve = s
	e.POST(apiUrl+"/apps", middlewares.AuthorizeClientRequest(), createApp)
	e.GET(apiUrl+"/apps", middlewares.AuthorizeClientRequest(), fetchApps)
	e.GET(apiUrl+"/apps/:id", middlewares.AuthorizeClientRequest(), GetAppById)
	e.PUT(apiUrl+"/apps", middlewares.AuthorizeClientRequest(), updateApp)
}

func createApp(c *gin.Context) {
	var req core.CreateAppRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	response := serve.AppService.CreateApp(req)

	if response.Error {
		c.JSON(response.Code, gin.H{
			"message": response.Meta.Message,
		})
		return
	}

	c.JSON(response.Code, response.Meta)
}

func updateApp(c *gin.Context) {
	var req core.UpdateAppRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	response := serve.AppService.UpdateApp(req)

	if response.Error {
		c.JSON(response.Code, gin.H{
			"message": response.Meta.Message,
		})
		return
	}

	c.JSON(response.Code, response.Meta)
}

func fetchApps(c *gin.Context) {

	response := serve.AppService.FetchApps()
	if response.Error {
		c.JSON(response.Code, gin.H{
			"message": response.Meta.Message,
		})
		return
	}

	c.JSON(response.Code, response.Meta)
}

func GetAppById(c *gin.Context) {
	idStr := c.Param("id")
	response := serve.AppService.GetApp(idStr)
	if response.Error {
		c.JSON(response.Code, gin.H{
			"message": response.Meta.Message,
		})
		return
	}

	c.JSON(response.Code, response.Meta)
}
