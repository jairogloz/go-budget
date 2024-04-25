package account

import "github.com/jairogloz/go-budget/pkg/domain/core"

// Service implements core.AccountService and holds the required components to
// perform the operations related to the account domain.
type Service struct {
	repo core.AccountRepository
}

// NewService creates a new account service.
func NewService(repo core.AccountRepository) core.AccountService {
	return &Service{
		repo: repo,
	}
}
