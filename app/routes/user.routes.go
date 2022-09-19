package routes

import (
	//	"postoffice/app/models"
	"net/http"
	"postoffice/app/middlewares"
	"postoffice/app/models"
	"postoffice/app/services"

	"github.com/gin-gonic/gin"

	//	"postoffice/app/utils"
	"postoffice/app/core"
	//	"strconv"
	//	log "github.com/sirupsen/logrus"
)

func RegisterUserRoutes(e *gin.Engine, s services.Services) {

	e.POST(apiUrl+"/users", middlewares.AuthorizeClientRequest(), createUser)

	e.GET(apiUrl+"/users", middlewares.AuthorizeClientRequest(), fetchUsers)

	e.GET(apiUrl+"/users/:id", middlewares.AuthorizeClientRequest(), getUser)

	e.PUT(apiUrl+"/users", middlewares.AuthorizeClientRequest(), updateUser)

	e.POST(apiUrl+"/users/auth", loginUser)
}

func createUser(c *gin.Context) {
	var req core.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	response := serve.UserService.CreateUser(req)

	if response.Error {
		c.JSON(response.Code, gin.H{
			"message": response.Meta.Message,
		})
		return
	}

	c.JSON(response.Code, response.Meta)
}

func fetchUsers(c *gin.Context) {

	response := serve.UserService.FetchUsers()

	if response.Error {
		c.JSON(response.Code, gin.H{
			"message": response.Meta.Message,
		})
		return
	}

	c.JSON(response.Code, response.Meta)
}

func getUser(c *gin.Context) {
	idStr := c.Param("id")
	response := serve.UserService.GetUser(idStr)
	if response.Error {
		c.JSON(response.Code, gin.H{
			"message": response.Meta.Message,
		})
		return
	}
	c.JSON(response.Code, response.Meta)
}

func updateUser(c *gin.Context) {
	var req models.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	response := serve.UserService.UpdateUser(req)

	if response.Error {
		c.JSON(response.Code, gin.H{
			"message": response.Meta.Message,
		})
		return
	}

	c.JSON(response.Code, response.Meta)
}

func loginUser(c *gin.Context) {
	var req core.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	response := serve.UserService.Login(req)

	if response.Error {
		c.JSON(response.Code, gin.H{
			"message": response.Meta.Message,
		})
		return
	}

	c.JSON(response.Code, response.Meta)
}
