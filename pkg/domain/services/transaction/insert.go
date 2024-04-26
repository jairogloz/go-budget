package transaction

import "github.com/jairogloz/go-budget/pkg/domain/core"

// Insert inserts a new transaction into the database.
func (s Service) Insert(transaction *core.Transaction, newCategory bool) (*core.Account, error) {

	if transaction.Category != nil && newCategory {
		category := core.Category{
			Name:   *transaction.Category,
			UserId: transaction.UserId,
		}
		_, err := s.categoryRepo.Insert(&category)
		if err != nil {
			return nil, err
		}
	}

	return s.txRepo.Insert(transaction, false)
}
