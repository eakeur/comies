package order

import (
	"comies/app/core/types"
	"context"
)

type Actions interface {
	Create(ctx context.Context, o Order) error
	List(ctx context.Context, f Filter) ([]Order, error)
	GetByID(ctx context.Context, id types.ID) (Order, error)
	GetByCustomerPhone(ctx context.Context, phone string) (Order, error)
}
