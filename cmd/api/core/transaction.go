package core

import (
	"github.com/gin-gonic/gin"
)

// TransactionCreate is the struct that represents the request to create a new
// transaction.
type TransactionCreate struct {
	Amount      float64 `json:"amount" binding:"required"`
	AccountId   *string `json:"account_id"`
	Category    *string `json:"category"`
	Description string  `json:"description" bson:"description"`
}

// TransactionCategory is the struct that represents the category of a
// transaction.
type TransactionCategory struct {
	Name  string `json:"name" binding:"required"`
	IsNew *bool  `json:"is_new" binding:"required"`
}

// TransactionHandler exposes the handlers for the transactions services.
type TransactionHandler interface {
	Delete(c *gin.Context)
	Insert(c *gin.Context)
}
