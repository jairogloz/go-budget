package core

import (
	"github.com/gin-gonic/gin"
	"github.com/jairogloz/go-budget/pkg/domain/core"
	"time"
)

// TransactionCreate is the struct that represents the request to create a new
// transaction.
type TransactionCreate struct {
	Amount      float64              `json:"amount" binding:"required"`
	AccountId   string               `json:"account_id" binding:"required"`
	Category    *TransactionCategory `json:"category"`
	Description string               `json:"description" bson:"description"`
}

// TransactionCreateResponse is the struct that represents the response to a
// transaction creation request.
type TransactionCreateResponse struct {
	Transaction core.Transaction `json:"transaction"`
	Account     core.Account     `json:"account"`
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

// ToDomain converts a TransactionCreate into a core.Transaction.
func (t TransactionCreate) ToDomain(userId string) *core.Transaction {
	now := time.Now().UTC()

	domainT := &core.Transaction{
		Amount:      t.Amount,
		AccountId:   t.AccountId,
		Description: t.Description,
		CreatedAt:   &now,
		UpdatedAt:   &now,
		UserId:      userId,
	}

	if t.Category != nil {
		domainT.Category = &t.Category.Name
	}

	return domainT
}
