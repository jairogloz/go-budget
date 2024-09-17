package ports

import (
	"context"
	"github.com/jairogloz/go-budget/pkg/domain/core"
)

// CategoryRepository exposes the methods to interact with the category storage.
type CategoryRepository interface {
	Insert(ctx context.Context, category *core.Category) (*core.Category, error)
}
