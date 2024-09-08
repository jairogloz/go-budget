package core

import (
	"time"
)

// Transaction reflects a transaction in the database. The amount could be
// positive or negative, depending on the transaction type.
type Transaction struct {
	Amount      float64    `json:"amount" bson:"amount"`
	AccountId   *string    `json:"account_id" bson:"account_id"`
	Category    *string    `json:"category" bson:"category"`
	CreatedAt   *time.Time `json:"created_at" bson:"created_at"`
	Description string     `json:"description" bson:"description"`
	ID          string     `json:"id" bson:"_id"`
	SubCategory *string    `json:"sub_category" bson:"sub_category"`
	UpdatedAt   *time.Time `json:"updated_at" bson:"updated_at"`
	UserId      string     `json:"user_id" bson:"user_id"`
}

// TransactionCreateParams represents the parameters that can be used to create
// a new transaction.
type TransactionCreateParams struct {
	AccountID   *string `json:"account_id" bson:"account_id"`
	Amount      float64 `json:"amount" bson:"amount"`
	Category    *string `json:"category" bson:"category"`
	Description string  `json:"description" bson:"description"`
	SubCategory *string `json:"sub_category" bson:"sub_category"`
}

// TransactionUpdateParams represents the parameters that can be updated in a
// transaction.
//
// Fields that are not nil will be updated in the database, while nil fields
// will be ignored.
type TransactionUpdateParams struct {
	Amount      *float64   `json:"amount" bson:"amount"`
	AccountID   *string    `json:"account_id" bson:"account_id"`
	Category    *string    `json:"category" bson:"category"`
	Description *string    `json:"description" bson:"description"`
	UpdatedAt   *time.Time `json:"updated_at" bson:"updated_at"`
}
