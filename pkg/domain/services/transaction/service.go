package transaction

import (
	"github.com/jairogloz/go-budget/pkg/domain/ports"
)

// Service implements core.TransactionService and holds the required components to
// perform the operations related to the transactions domain.
type Service struct {
	categoryRepo ports.CategoryRepository
	txRepo       ports.TransactionRepository
}

// NewService creates a new transaction service.
func NewService(txRepo ports.TransactionRepository,
	catRepo ports.CategoryRepository) ports.TransactionService {
	return &Service{
		categoryRepo: catRepo,
		txRepo:       txRepo,
	}
}
