package routes

import (
	"net/http"
	"postoffice/app/core"
	"postoffice/app/middlewares"
	"postoffice/app/services"

	"github.com/gin-gonic/gin"
)

func RegisterDomainRoutes(e *gin.Engine, s services.Services) {
	serve = s
	e.POST(apiUrl+"/domains", middlewares.AuthorizeClientRequest(), createDomain)
	e.GET(apiUrl+"/domains", middlewares.AuthorizeClientRequest(), fetchDomains)
	e.GET(apiUrl+"/domains/:id", middlewares.AuthorizeClientRequest(), getDomains)
	e.PUT(apiUrl+"/domains", middlewares.AuthorizeClientRequest(), updateDomain)
}

func createDomain(c *gin.Context) {
	var req core.CreateDomainRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	response := serve.DomainService.CreateDomain(req)

	if response.Error {
		c.JSON(response.Code, gin.H{
			"message": response.Meta.Message,
		})
		return
	}

	c.JSON(response.Code, response.Meta)
}

func updateDomain(c *gin.Context) {
	var req core.UpdateDomainRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	response := serve.DomainService.UpdateDomain(req)

	if response.Error {
		c.JSON(response.Code, gin.H{
			"message": response.Meta.Message,
		})
		return
	}

	c.JSON(response.Code, response.Meta)
}

func fetchDomains(c *gin.Context) {

	response := serve.DomainService.FetchDomains()
	if response.Error {
		c.JSON(response.Code, gin.H{
			"message": response.Meta.Message,
		})
		return
	}

	c.JSON(response.Code, response.Meta)
}

func getDomains(c *gin.Context) {

	idStr := c.Param("id")
	response := serve.DomainService.GetDomain(idStr)
	if response.Error {
		c.JSON(response.Code, gin.H{
			"message": response.Meta.Message,
		})
		return
	}

	c.JSON(response.Code, response.Meta)
}
