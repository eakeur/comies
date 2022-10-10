package ordering

import (
	"comies/app/core/id"
	"comies/app/core/reservation"
	"context"
)

//go:generate moq -fmt goimports -out menu_service_mock.go . MenuService:MenuServiceMock

type (
	MenuService interface {
		Reserve(ctx context.Context, reservation reservation.Reservation) ([]reservation.Failure, error)
		UpdateReservation(ctx context.Context, reservationID id.ID, consume bool) error
	}
)
