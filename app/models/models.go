package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Issue - struct to map with mongodb documents
type Issue struct {
	ID          primitive.ObjectID `bson:"_id"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	Title       string             `bson:"title"`
	Code        string             `bson:"code"`
	Description string             `bson:"description"`
	Completed   bool               `bson:"completed"`
}

type User struct {
	ID          primitive.ObjectID `bson:"_id"`
	CreatedAt   time.Time          `bson:"created_at"`
	Action      string             `bson:"action"`
	Email       string             `bson:"email"`
	UserId      string             `bson:"code"`
	Description string             `bson:"description"`
}
