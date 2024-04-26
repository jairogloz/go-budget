package transaction

import "github.com/jairogloz/go-budget/pkg/domain/core"

// Insert inserts a new transaction into the database.
func (s Service) Insert(transaction *core.Transaction, newCategory bool) (*core.Account, error) {
	return s.repo.Insert(transaction, false)
}
