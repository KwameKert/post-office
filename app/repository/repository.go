package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	Issues *issueLayer
	Users  *userLayer
	Apps   *appLayer
}

func NewRepository(db *mongo.Database) Repo {
	return Repo{
		Issues: newIssueRepoLayer(db),
		Users:  newUserRepoLayer(db),
		Apps:   newAppRepoLayer(db),
		//	Tasks: newTaskRepoLayer(db),
		// Transactions:      newTransactionLayer(db),
		// Wallets:           newWalletLayer(db),
		// TransactionEvents: newEventLayer(db),
	}

}
