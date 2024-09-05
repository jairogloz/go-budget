package ports

import "github.com/jairogloz/go-budget/pkg/domain/core"

// CategoryRepository exposes the methods to interact with the category storage.
type CategoryRepository interface {
	Insert(category *core.Category) (*core.Category, error)
}
