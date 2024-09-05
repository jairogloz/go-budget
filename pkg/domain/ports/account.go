package ports

import "github.com/jairogloz/go-budget/pkg/domain/core"

// AccountRepository exposes the methods to interact with the account storage.
type AccountRepository interface {
	Create(user *core.User, account *core.Account) error
	CountAccounts(userId string) (int, error)
	Delete(userId, id string) error
	GetByID(userId, id string) (core.Account, error)
	List(userId string) ([]core.Account, error)
}

// AccountService exposes the services provided by this application on the account domain.
type AccountService interface {
	Create(user *core.User, account *core.Account) error
	Delete(userId, id string) error
	GetByID(userId, id string) (core.Account, error)
	List(userId string) ([]core.Account, error)
	Update(userID string, accountID string, updateParams core.AccountUpdateParams) error
}