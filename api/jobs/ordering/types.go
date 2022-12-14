package ordering

import (
	"comies/core/types"
	"context"
)

type PriceFetcher func(ctx context.Context, productID types.ID) (types.Currency, error)

type Dispatcher func(ctx context.Context, d types.Dispatcher) error

type BillCreator func(ctx context.Context) error
