package product

import (
	"context"
	"gomies/app/sdk/types"
)

//go:generate moq -fmt goimports -out stock_service_mock.go . StockService:StockServiceMock

type (
	StockService interface {
		Compute(ctx context.Context, productID types.ID) (types.Quantity, error)
		ComputeSome(ctx context.Context, productID ...types.ID) ([]types.Quantity, error)
	}
)
