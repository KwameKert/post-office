package routes

import (
	//	"postoffice/app/models"
	"net/http"
	"postoffice/app/services"

	"github.com/gin-gonic/gin"

	//	"postoffice/app/utils"
	"postoffice/app/core"
	//	"strconv"
	//	log "github.com/sirupsen/logrus"
)

func RegisterUserRoutes(e *gin.Engine, s services.Services) {

	e.POST("/users", createUser)

	e.GET("/users", fetchUsers)

	e.GET("/users/:id", getUser)

	// e.PUT("/users", func(c *gin.Context) {
	// 	var req models.User

	// 	if err := c.ShouldBindJSON(&req); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"message": err.Error(),
	// 		})
	// 		return
	// 	}
	// 	response := s.UserService.UpdateUser(req)

	// 	if response.Error {
	// 		c.JSON(response.Code, gin.H{
	// 			"message": response.Meta.Message,
	// 		})
	// 		return
	// 	}

	// 	c.JSON(response.Code, response.Meta)
	// })

	// e.GET("/users", func(c *gin.Context) {

	// 	response := s.UserService.FetchUsers()

	// 	if response.Error {
	// 		c.JSON(response.Code, gin.H{
	// 			"message": response.Meta.Message,
	// 		})
	// 		return
	// 	}

	// 	c.JSON(response.Code, response.Meta)
	// })

	// e.GET("/users/:id", func(c *gin.Context) {
	// 	idStr := c.Param("id")
	// 	id, _ := strconv.Atoi(idStr)
	// 	response := s.UserService.GetUser(id)
	// 	if response.Error {
	// 		c.JSON(response.Code, gin.H{
	// 			"message": response.Meta.Message,
	// 		})
	// 		return
	// 	}
	// 	c.JSON(response.Code, response.Meta)
	// })

	// e.DELETE("/users/:id", func(c *gin.Context) {
	// 	idStr := c.Param("id")
	// 	id, _ := strconv.Atoi(idStr)
	// 	response := s.UserService.DeleteUser(id)
	// 	if response.Error {
	// 		c.JSON(response.Code, gin.H{
	// 			"message": response.Meta.Message,
	// 		})
	// 		return
	// 	}
	// 	c.JSON(response.Code, response.Meta)
	// })

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
