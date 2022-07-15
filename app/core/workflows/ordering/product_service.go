package ordering

import (
	"comies/app/sdk/types"
	"context"
)

//go:generate moq -fmt goimports -out product_service_mock.go . ProductService:ProductServiceMock

type (
	ProductService interface {
		ReserveResources(ctx context.Context, reservationID types.ID, reservation Reservation) (Reservation, error)
		UpdateResources(ctx context.Context, reservationID types.ID, consume bool) error
	}
)
