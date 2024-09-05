package core

import (
	"time"
)

// Transaction reflects a transaction in the database. The amount could be
// positive or negative, depending on the transaction type.
type Transaction struct {
	Amount      float64     `json:"amount" bson:"amount"`
	AccountId   *string     `json:"account_id" bson:"account_id"`
	Category    *string     `json:"category" bson:"category"`
	CreatedAt   *time.Time  `json:"created_at" bson:"created_at"`
	Description string      `json:"description" bson:"description"`
	ID          interface{} `json:"id" bson:"_id"`
	UpdatedAt   *time.Time  `json:"updated_at" bson:"updated_at"`
	UserId      string      `json:"user_id" bson:"user_id"`
}
