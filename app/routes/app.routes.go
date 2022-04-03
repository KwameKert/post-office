package routes

import (
	"net/http"
	"postoffice/app/core"
	"postoffice/app/services"

	"github.com/gin-gonic/gin"
)

func RegisterAppRoutes(e *gin.Engine, s services.Services) {
	serve = s
	e.POST(apiUrl+"/apps", createApp)
	e.GET(apiUrl+"/apps", fetchApps)

	// e.PUT(apiUrl+"/users", updateUser)

	// e.GET(apiUrl+"/users", middleware.AuthorizeClientRequest(), fetchUsers)

	// e.GET(apiUrl+"/users/:id", middleware.AuthorizeClientRequest(), getUserById)

	// e.DELETE(apiUrl+"/users/:id", middleware.AuthorizeClientRequest(), deleteUserById)
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
