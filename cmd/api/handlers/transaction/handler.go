package transaction

import (
	ginCore "github.com/jairogloz/go-budget/cmd/api/core"
	domainCore "github.com/jairogloz/go-budget/pkg/domain/core"
)

// Handler implements the core.TransactionHandler interface.
type Handler struct {
	service domainCore.TransactionService
}

// NewHandler creates a new transaction handler.
func NewHandler(service domainCore.TransactionService) ginCore.TransactionHandler {
	return &Handler{
		service: service,
	}
}
