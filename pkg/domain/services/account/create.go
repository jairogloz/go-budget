package account

import (
	"github.com/jairogloz/go-budget/pkg/domain/core"
	"log"
)

// Create creates a new account.
func (s Service) Create(account *core.Account) error {
	log.Println("creating account")
	return s.repo.Create(account)
}
