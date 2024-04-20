package core

import (
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type Account struct {
	ID             string          `json:"id" bson:"_id"`
	Name           string          `json:"name" bson:"name"`
	CurrentBalance decimal.Decimal `json:"current_balance" bson:"current_balance"`
}

// AccountHandlers exposes the handlers for the account domain.
type AccountHandlers struct {
	CreateAccount func(c *gin.Context)
}
