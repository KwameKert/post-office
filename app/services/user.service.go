package services

import (
	//	"errors"
	"postoffice/app/core"
	"postoffice/app/models"
	"postoffice/app/pkg"
	"postoffice/app/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	//	"gorm.io/gorm"
)

type userServiceLayer struct {
	repository repository.Repo
	config     *core.Config
}

func newUserServiceLayer(r repository.Repo, c *core.Config) *userServiceLayer {
	return &userServiceLayer{
		repository: r,
		config:     c,
	}
}

func (u *userServiceLayer) CreateUser(req core.CreateUserRequest) core.Response {
	user := models.User{
		Username:  req.Username,
		Password:  pkg.HashPassword(req.Password),
		Email:     req.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := u.repository.Users.Create(&user); err != nil {
		return core.Error(err, nil)
	}

	return core.Success(&map[string]interface{}{
		"user": user,
	}, core.String("user created successfully"))
}

func (u *userServiceLayer) FetchUsers() core.Response {
	err, users := u.repository.Users.Fetch()
	if err != nil {
		return core.Error(err, nil)
	}
	if len(users) == 0 {
		return core.NoContentFound(err, core.String("No users found"))
	}

	return core.Success(&map[string]interface{}{
		"users": users,
	}, core.String("users found successfully"))
}

func (u *userServiceLayer) GetUser(id string) core.Response {
	user := models.User{}
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return core.Error(err, nil)
	}

	if err := u.repository.Users.Get(&user, objectId); err != nil {
		return core.BadRequest(err, nil)
	}

	return core.Success(&map[string]interface{}{
		"user": user,
	}, core.String("user found successfully"))
}

// func (u *userServiceLayer) DeleteUser(id int) core.Response {
// 	user := models.User{}

// 	if err := u.repository.Users.Get(&user, id); err != nil {
// 		return core.BadRequest(err, nil)
// 	}

// 	if err := u.repository.Users.Delete(&user, id); err != nil {
// 		return core.BadRequest(err, nil)
// 	}

// 	return core.Success(&map[string]interface{}{}, core.String("user deleted successfully"))
// }

func (u *userServiceLayer) UpdateUser(user models.User) core.Response {
	userDTO := models.User{}

	if err := u.repository.Users.Get(&userDTO, user.Id); err != nil {
		return core.BadRequest(err, nil)
	}
	if err := u.repository.Users.Update(&user); err != nil {
		return core.Error(err, nil)
	}

	return core.Success(&map[string]interface{}{
		"user": user,
	}, core.String("users updated successfully"))
}
