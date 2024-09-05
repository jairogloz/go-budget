package account

import (
	"github.com/jairogloz/go-budget/pkg/domain/ports"
)

// Service implements core.AccountService and holds the required components to
// perform the operations related to the account domain.
type Service struct {
	repo ports.AccountRepository
}

// NewService creates a new account service.
func NewService(repo ports.AccountRepository) ports.AccountService {
	return &Service{
		repo: repo,
	}
}
