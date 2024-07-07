package core

import "time"

// Account reflects a monetary account in the system.
type Account struct {
	Balance        float64     `json:"balance" bson:"balance"`
	CreatedAt      *time.Time  `json:"created_at" bson:"created_at"`
	ID             interface{} `json:"id" bson:"_id"`
	InitialBalance float64     `json:"initial_balance" bson:"initial_balance"`
	Name           string      `json:"name" bson:"name"`
	UpdatedAt      *time.Time  `json:"updated_at" bson:"updated_at"`
	UserId         string      `json:"user_id" bson:"user_id"`
}

// AccountRepository exposes the methods to interact with the account storage.
type AccountRepository interface {
	Create(user *User, account *Account) error
	CountAccounts(userId string) (int, error)
	Delete(userId, id string) error
	GetByID(userId, id string) (Account, error)
	List(userId string) ([]Account, error)
}

// AccountService exposes the services provided by this application on the account domain.
type AccountService interface {
	Create(user *User, account *Account) error
	Delete(userId, id string) error
	GetByID(userId, id string) (Account, error)
	List(userId string) ([]Account, error)
}
