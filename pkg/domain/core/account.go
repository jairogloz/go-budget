package core

// Account reflects a monetary account in the system.
type Account struct {
	Balance float64 `json:"balance" bson:"balance"`
	ID      string  `json:"id" bson:"_id"`
	Name    string  `json:"name" bson:"name"`
	UserId  string  `json:"user_id" bson:"user_id"`
}

// AccountRepository exposes the methods to interact with the account storage.
type AccountRepository interface {
	GetByID(userId, id string) (Account, error)
	List(userId string) ([]Account, error)
}

// AccountService exposes the services provided by this application on the account domain.
type AccountService interface {
	GetByID(userId, id string) (Account, error)
	List(userId string) ([]Account, error)
}
