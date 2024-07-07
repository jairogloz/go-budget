package account

import (
	ginCore "github.com/jairogloz/go-budget/cmd/gin/core"
	domainCore "github.com/jairogloz/go-budget/pkg/domain/core"
	"github.com/jairogloz/go-budget/pkg/domain/ports"
)

// Handler implements the core.AccountHandler interface.
type Handler struct {
	service domainCore.AccountService
	ctxHdl  ports.ContextHandler
}

// NewHandler creates a new account handler.
func NewHandler(service domainCore.AccountService, ctxHdl ports.ContextHandler) ginCore.AccountHandler {
	return &Handler{
		ctxHdl:  ctxHdl,
		service: service,
	}
}
