package transaction

import (
	"fmt"
	"github.com/jairogloz/go-budget/pkg/domain/core"
	"time"
)

// Insert inserts a new transaction into the database.
//
// It takes a user and transaction creation parameters as input, constructs a new
// transaction object, and inserts it into the database. If the insertion is successful,
// it returns the newly created transaction. If there is an error during insertion, it
// returns an error.
//
// Parameters:
//   - user: the user who is creating the transaction
//   - transactionCreateParams: the parameters required to create a new transaction
//
// Returns:
//   - newTx: the newly created transaction
//   - err: an error if there is an issue during insertion
func (s Service) Insert(user *core.User, transactionCreateParams core.TransactionCreateParams) (newTx *core.Transaction, err error) {

	now := time.Now().UTC()
	tx := &core.Transaction{
		Amount:      transactionCreateParams.Amount,
		AccountId:   transactionCreateParams.AccountID,
		Category:    transactionCreateParams.Category,
		CreatedAt:   &now,
		Description: transactionCreateParams.Description,
		SubCategory: transactionCreateParams.SubCategory,
		UpdatedAt:   &now,
		UserId:      user.ID,
	}

	insertedID, err := s.txRepo.Insert(tx)
	if err != nil {
		return nil, fmt.Errorf("error inserting transaction: %w", err)
	}
	tx.ID = insertedID

	return tx, nil
}
