package product

import (
	"context"
	"gomies/app/core/entities/catalog/product"
	"gomies/pkg/sdk/types"
	"time"
)

//go:generate moq -fmt goimports -out stock_service_mock.go . StockService:StockServiceMock

type (
	StockService interface {
		Compute(ctx context.Context, productID types.UID) (types.Quantity, error)
		ComputeSome(ctx context.Context, productID ...types.UID) ([]types.Quantity, error)
		CreateMovement(ctx context.Context, config product.Stock, resourceID types.UID, movement Movement) (types.Quantity, error)
	}

	Movement struct {
		ProductID      types.UID
		Output         bool
		Date           time.Time
		Quantity       types.Quantity
		PaidValue      types.Currency
		Agent          types.UID
		Batch          string
		ShelfLife      time.Time
		AdditionalData string
	}
)
