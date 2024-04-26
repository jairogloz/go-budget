package account

import "github.com/jairogloz/go-budget/pkg/domain/core"

func (s Service) GetByID(userId, id string) (core.Account, error) {
	return s.repo.GetByID(userId, id)
}
