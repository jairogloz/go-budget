package account

import "github.com/jairogloz/go-budget/pkg/domain/core"

// List returns the list of accounts for a given user.
func (s Service) List(userId string) ([]core.Account, error) {

	return s.repo.List(userId)
}
