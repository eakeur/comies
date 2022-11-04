package price

import (
	"comies/app/core/types"
	"context"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	Create(ctx context.Context, p Price) error
	List(ctx context.Context, productID types.ID) ([]Price, error)
	GetLatestValue(ctx context.Context, productID types.ID) (types.Currency, error)
}
