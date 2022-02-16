package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	Issues *issueLayer
	Users  *userLayer
}

func NewRepository(db *mongo.Database) Repo {
	return Repo{
		Issues: newIssueRepoLayer(db),
		Users:  newUserRepoLayer(db),
		//	Tasks: newTaskRepoLayer(db),
		// Transactions:      newTransactionLayer(db),
		// Wallets:           newWalletLayer(db),
		// TransactionEvents: newEventLayer(db),
	}

}
