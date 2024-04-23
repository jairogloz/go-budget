package core

import (
	"github.com/shopspring/decimal"
	"time"
)

// Transaction reflects a transaction in the database. The amount could be
// positive or negative, depending on the transaction type.
type Transaction struct {
	Amount      decimal.Decimal `json:"amount" bson:"amount"`
	AccountId   string          `json:"account_id" bson:"account_id"`
	Category    *string         `json:"category" bson:"category"`
	CreatedAt   *time.Time      `json:"created_at" bson:"created_at"`
	Description string          `json:"description" bson:"description"`
	ID          string          `json:"id" bson:"_id"`
	UpdatedAt   *time.Time      `json:"updated_at" bson:"updated_at"`
}

// TransactionRepository exposes the methods to interact with the transaction
// storage.
type TransactionRepository interface {
	//FindByAccountID(accountID string) ([]Transaction, error)
	Insert(transaction *Transaction, newCategory bool) error
}

// TransactionService exposes the services provided by this application on the
// transaction domain.
type TransactionService interface {
	Insert(transaction *Transaction, newCategory bool) error
}
