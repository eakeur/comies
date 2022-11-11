package product

import (
	"comies/app/core/types"
	"context"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	List(ctx context.Context, filter Filter) ([]Product, error)
	GetByID(ctx context.Context, id types.ID) (Product, error)
	GetNameByID(ctx context.Context, id types.ID) (string, error)
	Create(ctx context.Context, p Product) error
	Update(ctx context.Context, p Product) error
	Remove(ctx context.Context, id types.ID) error
}
