package ports

import (
	"context"
	domainCore "github.com/jairogloz/go-budget/pkg/domain/core"
)

// ContextHandler exposes the methods to handle the context.
type ContextHandler interface {
	GetUser(c context.Context) (*domainCore.User, error)
}
