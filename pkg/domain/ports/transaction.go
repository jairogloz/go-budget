package ports

import (
	"context"
	"github.com/jairogloz/go-budget/pkg/domain/core"
	"time"
)

// TransactionRepository exposes the methods to interact with the transaction
// storage.
type TransactionRepository interface {
	Delete(userId, transactionID string) error
	FindByAccountID(userId, accountID string) ([]core.Transaction, error)
	Insert(transaction *core.Transaction) (insertedID string, err error)
	List(ctx context.Context, userID string, from, to *time.Time, limit, offset int) (core.Transactions, error)
}

// TransactionService exposes the services provided by this application on the
// transaction domain.
type TransactionService interface {
	Delete(userId, transactionID string) error
	FindByAccountID(userId, accountID string) ([]core.Transaction, error)
	Insert(user *core.User, transactionCreateParams core.TransactionCreateParams) (newTx *core.Transaction, err error)
	List(ctx context.Context, user *core.User, from, to, limit, offset, listType string) (result interface{}, err error)
}
