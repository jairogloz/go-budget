package transaction

import (
	ginCore "github.com/jairogloz/go-budget/cmd/api/core"
	"github.com/jairogloz/go-budget/pkg/domain/ports"
)

// Handler implements the core.TransactionHandler interface.
type Handler struct {
	service ports.TransactionService
}

// NewHandler creates a new transaction handler.
func NewHandler(service ports.TransactionService) ginCore.TransactionHandler {
	return &Handler{
		service: service,
	}
}
