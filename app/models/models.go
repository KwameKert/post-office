package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

//Issue - struct to map with mongodb documents

type App struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Status      string             `bson:"status" json:"status"`
	Modules     []Module           `bson:"modules" json:"modules"`
}

type Module struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Status      string             `bson:"status" json:"status"`
	Domains     []Domain           `bson:"domains" json:"domains"`
	AppID       primitive.ObjectID `bson:"app_id" json:"app_id"`
}

type Domain struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Status      string             `bson:"status" json:"status"`
	Logs        []Log              `bson:"logs" json:"logs"`
	ModuleId    primitive.ObjectID `bson:"module_id" json:"module_id"`
}

type Log struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	Data      string             `bson:"data" json:"data"`
	DomainId  primitive.ObjectID `bson:"domain_id" json:"domain_id"`
	Action    string             `bson:"action"`
	Creator   string             `bson:"user_id"`
}

type User struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
	Username    string             `json:"username" bson:"username"`
	Email       string             `json:"email" bson:"email"`
	Password    string             `json:"password" bson:"password"`
	Description string             `json:"description" bson:"description"`
}
