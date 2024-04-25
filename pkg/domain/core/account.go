package core

// Account reflects a monetary account in the system.
type Account struct {
	CurrentBalance float64 `json:"current_balance" bson:"current_balance"`
	ID             string  `json:"id" bson:"_id"`
	Name           string  `json:"name" bson:"name"`
	UserId         string  `json:"user_id" bson:"user_id"`
}

// AccountRepository exposes the methods to interact with the account storage.
type AccountRepository interface {
	List(userId string) ([]Account, error)
}

// AccountService exposes the services provided by this application on the account domain.
type AccountService interface {
	List(userId string) ([]Account, error)
}
