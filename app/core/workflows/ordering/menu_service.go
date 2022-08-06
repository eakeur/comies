package ordering

import (
	"comies/app/core/entities/reservation"
	"comies/app/core/types"
	"context"
)

//go:generate moq -fmt goimports -out menu_service_mock.go . MenuService:MenuServiceMock

type (
	MenuService interface {
		Reserve(ctx context.Context, reservation reservation.Reservation) ([]reservation.Failure, error)
		UpdateReservation(ctx context.Context, reservationID types.ID, consume bool) error
	}
)
