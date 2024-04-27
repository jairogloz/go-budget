package core

import (
	"github.com/gin-gonic/gin"
	"github.com/jairogloz/go-budget/pkg/domain/core"
	"time"
)

// AccountCreate is the struct that represents the request to create a new account.
type AccountCreate struct {
	Name           string   `json:"name" binding:"required"`
	InitialBalance *float64 `json:"initial_balance" binding:"required"`
}

// ToDomain converts an AccountCreate into a core.Account.
func (a AccountCreate) ToDomain(userId string) *core.Account {
	now := time.Now().UTC()
	return &core.Account{
		Name:           a.Name,
		InitialBalance: *a.InitialBalance,
		Balance:        *a.InitialBalance,
		UserId:         userId,
		CreatedAt:      &now,
		UpdatedAt:      &now,
	}
}

// AccountHandler exposes the handlers for the account domain.
type AccountHandler interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	List(c *gin.Context)
	GetById(c *gin.Context)
}
