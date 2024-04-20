package core

import "github.com/shopspring/decimal"

type Transaction struct {
	Amount      decimal.Decimal `json:"amount" bson:"amount"`
	Category    *string         `json:"category" bson:"category"`
	Description string          `json:"description" bson:"description"`
	ID          string          `json:"id" bson:"_id"`
}
