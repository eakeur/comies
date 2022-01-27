package category

import (
	"context"
	"gomies/app/core/types/id"
)

type Actions interface {
	Get(context.Context, id.External) (Category, error)
	List(context.Context, Filter) ([]Category, error)
	Create(context.Context, Category) (Category, error)
	Remove(context.Context, id.External) error
}

