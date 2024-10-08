package account

import (
	"fmt"
	"github.com/jairogloz/go-budget/pkg/domain/core"
)

// Create creates a new account.
func (s Service) Create(user *core.User, account core.Account) (*core.Account, error) {
	accountCount, err := s.repo.CountAccounts(user.ID)
	if err != nil {
		// Todo: log error
		return nil, fmt.Errorf("error counting accounts: %w", err)
	}
	if accountCount >= user.FeatureAccess.MaxAccounts {
		return nil, fmt.Errorf("account limit '%d' reached for user level %s", user.FeatureAccess.MaxAccounts,
			user.Level)
	}
	insertedAccountID, err := s.repo.Create(account)
	if err != nil {
		// Todo: log error
		return nil, fmt.Errorf("error creating account: %w", err)
	}
	account.ID = insertedAccountID

	return &account, nil
}
