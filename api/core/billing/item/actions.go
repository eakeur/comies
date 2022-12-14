package item

import (
	"comies/core/types"
	"context"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	SumByBillID(ctx context.Context, billID types.ID) (types.Currency, error)
	List(ctx context.Context, filter Filter) ([]Item, error)
	Create(ctx context.Context, m Item) error
}
