package core

import "time"

// Account reflects a monetary account in the system.
type Account struct {
	Balance        float64    `json:"balance" bson:"balance"`
	CreatedAt      *time.Time `json:"created_at" bson:"created_at"`
	ID             string     `json:"id" bson:"_id"`
	InitialBalance float64    `json:"initial_balance" bson:"initial_balance"`
	Name           string     `json:"name" bson:"name"`
	UpdatedAt      *time.Time `json:"updated_at" bson:"updated_at"`
	UserId         string     `json:"user_id" bson:"user_id"`
}

type AccountUpdateParams struct {
	Name string `json:"name" bson:"name,omitempty"`
}
