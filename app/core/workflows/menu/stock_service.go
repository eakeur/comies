package menu

import (
	"comies/app/core/entities/ingredient"
	"comies/app/sdk/types"
	"context"
)

//go:generate moq -fmt goimports -out stock_service_mock.go . StockService:StockServiceMock

type (
	StockService interface {
		ReserveResources(ctx context.Context, reservationID types.ID, resources ...ingredient.Ingredient) ([]ItemFailed, error)
		ConsumeResources(ctx context.Context, reservationID types.ID) error
		FreeResources(ctx context.Context, reservationID types.ID) error
	}
)
