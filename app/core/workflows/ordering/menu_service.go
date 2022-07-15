package ordering

import (
	"comies/app/sdk/types"
	"context"
)

//go:generate moq -fmt goimports -out menu_service_mock.go . MenuService:MenuServiceMock

type (
	MenuService interface {
		ReserveResources(ctx context.Context, reservationID types.ID, reservation Reservation) (Reservation, error)
		UpdateResources(ctx context.Context, reservationID types.ID, consume bool) error
	}
)
