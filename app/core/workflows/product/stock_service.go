package product

import (
	"context"
	"gomies/app/core/entities/catalog/product"
	"gomies/app/sdk/types"
)

//go:generate moq -fmt goimports -out stock_service_mock.go . StockService:StockServiceMock

type (
	StockService interface {
		ReserveResources(ctx context.Context, reservationID types.ID, resources ...product.Ingredient) ([]FailedReservation, error)
		ConsumeResources(ctx context.Context, reservationID types.ID) error
		FreeResources(ctx context.Context, reservationID types.ID) error
	}
)
