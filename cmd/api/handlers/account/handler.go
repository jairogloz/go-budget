package account

import (
	ginCore "github.com/jairogloz/go-budget/cmd/api/core"
	"github.com/jairogloz/go-budget/pkg/domain/ports"
)

// Handler implements the core.AccountHandler interface.
type Handler struct {
	service ports.AccountService
	ctxHdl  ports.ContextHandler
}

// NewHandler creates a new account handler.
func NewHandler(service ports.AccountService, ctxHdl ports.ContextHandler) ginCore.AccountHandler {
	return &Handler{
		ctxHdl:  ctxHdl,
		service: service,
	}
}
