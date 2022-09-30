package routes

import (
	"net/http"
	"postoffice/app/core"
	"postoffice/app/services"

	"github.com/gin-gonic/gin"
)

func RegisterLogRoutes(e *gin.Engine, s services.Services) {
	serve = s
	e.POST(apiUrl+"/logs", createLog)
	e.GET(apiUrl+"/logs", searchLogs)
}

func createLog(c *gin.Context) {
	var req core.CreateLogRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	response := serve.LogService.CreateLog(req)

	if response.Error {
		c.JSON(response.Code, gin.H{
			"message": response.Meta.Message,
		})
		return
	}
	c.JSON(response.Code, response.Meta)
}

func searchLogs(c *gin.Context) {

	response := serve.LogService.SearchLog()

	if response.Error {
		c.JSON(response.Code, gin.H{
			"message": response.Meta.Message,
		})
		return
	}
	c.JSON(response.Code, response.Meta)
}
