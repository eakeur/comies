package category

import (
	"context"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type (
	Actions interface {
		GetCategory(ctx context.Context, categoryKey Key) (Category, error)
		ListCategories(ctx context.Context, categoryFilter Filter) ([]Category, error)
		CreateCategory(ctx context.Context, cat Category) (Category, error)
		RemoveCategory(ctx context.Context, categoryID Key) error
		UpdateCategory(ctx context.Context, category Category) error
	}
)
