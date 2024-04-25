package account

import (
	ginCore "github.com/jairogloz/go-budget/cmd/gin/core"
	domainCore "github.com/jairogloz/go-budget/pkg/domain/core"
)

// Handler implements the core.AccountHandler interface.
type Handler struct {
	service domainCore.AccountService
}

// NewHandler creates a new account handler.
func NewHandler(service domainCore.AccountService) ginCore.AccountHandler {
	return &Handler{
		service: service,
	}
}
