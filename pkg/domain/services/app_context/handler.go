package app_context

import (
	"context"
	"fmt"
	"github.com/jairogloz/go-budget/pkg/domain/core"
	"github.com/jairogloz/go-budget/pkg/domain/ports"
)

var _ ports.ContextHandler = &Handler{}

type Handler struct {
}

// NewHandler creates a new context handler.
func NewHandler() ports.ContextHandler {
	return &Handler{}
}

// GetUser retrieves the user from the context.
func (h Handler) GetUser(c context.Context) (*core.User, error) {
	user, ok := c.Value(core.CtxKeyUser).(*core.User)
	if !ok {
		return nil, fmt.Errorf("user not found in the context")
	}

	return user, nil
}
