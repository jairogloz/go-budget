package transaction

import "github.com/jairogloz/go-budget/pkg/domain/core"

// Service implements core.TransactionService and holds the required components to
// perform the operations related to the transactions domain.
type Service struct {
	repo core.TransactionRepository
}

// NewService creates a new transaction service.
func NewService(repo core.TransactionRepository) core.TransactionService {
	return &Service{
		repo: repo,
	}
}
