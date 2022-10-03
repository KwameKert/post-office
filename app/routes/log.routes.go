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
	e.GET(apiUrl+"/logs/search", searchLogs)
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
	var queryData core.SearchRequest
	queryData.AppId = c.Query("appId")
	queryData.DomainId = c.Query("domainId")
	queryData.Text = c.Query("text")
	queryData.ModuleId = c.Query("moduleId")
	queryData.UserId = c.Query("userId")

	response := serve.LogService.SearchLog(queryData)

	if response.Error {
		c.JSON(response.Code, gin.H{
			"message": response.Meta.Message,
		})
		return
	}
	c.JSON(response.Code, response.Meta)
}
