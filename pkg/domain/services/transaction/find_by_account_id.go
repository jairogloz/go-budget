package transaction

import "github.com/jairogloz/go-budget/pkg/domain/core"

// FindByAccountID retrieves all the transactions for a given account and userId.
func (s Service) FindByAccountID(userId, accountID string) ([]core.Transaction, error) {
	return s.txRepo.FindByAccountID(userId, accountID)
}
