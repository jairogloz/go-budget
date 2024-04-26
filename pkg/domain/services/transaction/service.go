package transaction

import "github.com/jairogloz/go-budget/pkg/domain/core"

// Service implements core.TransactionService and holds the required components to
// perform the operations related to the transactions domain.
type Service struct {
	categoryRepo core.CategoryRepository
	txRepo       core.TransactionRepository
}

// NewService creates a new transaction service.
func NewService(txRepo core.TransactionRepository,
	catRepo core.CategoryRepository) core.TransactionService {
	return &Service{
		categoryRepo: catRepo,
		txRepo:       txRepo,
	}
}
