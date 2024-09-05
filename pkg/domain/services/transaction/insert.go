package transaction

import (
	"fmt"
	"github.com/jairogloz/go-budget/pkg/domain/core"
)

// Insert inserts a new transaction into the database.
func (s Service) Insert(transaction *core.Transaction) error {

	err := s.txRepo.Insert(transaction)
	if err != nil {
		return fmt.Errorf("error inserting transaction: %w", err)
	}

	return nil
}
