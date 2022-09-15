package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	Users *userLayer
	Apps  *appLayer
	Logs  *logLayer
}

func NewRepository(db *mongo.Database) Repo {
	return Repo{
		Users: newUserRepoLayer(db),
		Apps:  newAppRepoLayer(db),
		Logs:  newLogRepoLayer(db),
	}

}
