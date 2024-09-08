package ports

import (
	"github.com/jairogloz/go-budget/pkg/domain/core"
)

// TransactionRepository exposes the methods to interact with the transaction
// storage.
type TransactionRepository interface {
	Delete(userId, transactionID string) error
	FindByAccountID(userId, accountID string) ([]core.Transaction, error)
	Insert(transaction *core.Transaction) (insertedID string, err error)
}

// TransactionService exposes the services provided by this application on the
// transaction domain.
type TransactionService interface {
	Delete(userId, transactionID string) error
	FindByAccountID(userId, accountID string) ([]core.Transaction, error)
	Insert(user *core.User, transactionCreateParams core.TransactionCreateParams) (newTx *core.Transaction, err error)
}
