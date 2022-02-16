package services

import (
	//	"errors"
	"postoffice/app/core"
	"postoffice/app/models"
	"postoffice/app/repository"
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
		Email:       req.Email,
		Action:      "Nice",
		Description: req.Name,
	}
	if err := u.repository.Users.Create(&user); err != nil {
		return core.Error(err, nil)
	}

	return core.Success(&map[string]interface{}{
		"user": user,
	}, core.String("user logged successfully"))
}

// func (u *userServiceLayer) FetchUsers() core.Response {
// 	var users []models.User
// 	err := u.repository.Users.Fetch(&users)
// 	if err != nil {
// 		return core.Error(err, nil)
// 	}
// 	if len(users) < 1 {
// 		return core.NoContentFound(err, core.String("No users found"))
// 	}
//
// 	return core.Success(&map[string]interface{}{
// 		"users": users,
// 	}, core.String("users found successfully"))
// }

// func (u *userServiceLayer) GetUser(id int) core.Response {
// 	user := models.User{}

// 	if err := u.repository.Users.Get(&user, id); err != nil {
// 		return core.BadRequest(err, nil)
// 	}

// 	return core.Success(&map[string]interface{}{
// 		"user": user,
// 	}, core.String("users found successfully"))
// }

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

// func (u *userServiceLayer) UpdateUser(user models.User) core.Response {
// 	userDTO := models.User{}

// 	if err := u.repository.Users.Get(&userDTO, user.Id); err != nil {
// 		return core.BadRequest(err, nil)
// 	}
// 	userUpdated, err := u.repository.Users.Update(&user)
// 	if err != nil {
// 		return core.Error(err, nil)
// 	}

// 	return core.Success(&map[string]interface{}{
// 		"user": userUpdated,
// 	}, core.String("users updated successfully"))
// }
