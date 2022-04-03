package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

//Issue - struct to map with mongodb documents
type Issue struct {
	Model
	Title       string `bson:"title"`
	Code        string `bson:"code"`
	Description string `bson:"description"`
	Completed   bool   `bson:"completed"`
}

type App struct {
	Model
	Name        string `bson:"name"`
	Description string `bson:"description"`
	Status      string `bson:"status"`
}

type Log struct {
	Model,
	Data string `bson:"data"`
	AppID       int    `bson:"app_id"`
	Name        string `bson:"name"`
	Description string `bson:"description"`
}

type User struct {
	Model
	Action      string `bson:"action"`
	Email       string `bson:"email"`
	UserId      string `bson:"code"`
	Description string `bson:"description"`
}

type Model struct {
	//	Id        primitive.ObjectID `gorm:"primary_key" json:"id"`
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	//	DeletedAt *time.Time         `sql:"index" json:"deleted_at,omitempty"`
}
