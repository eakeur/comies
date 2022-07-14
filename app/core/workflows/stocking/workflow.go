package stocking

import (
	"context"
	"gomies/app/core/entities/movement"
	"gomies/app/core/entities/stock"
	"gomies/app/sdk/types"
)

var _ Workflow = workflow{}

//go:generate moq -fmt goimports -out workflow_mock.go . Workflow:WorkflowMock
type (
	Workflow interface {
		ConsumeResources(ctx context.Context, reservationID types.ID) error
		FreeResources(ctx context.Context, reservationID types.ID) error
		ReserveResources(ctx context.Context, reservationID types.ID, reservations []Reservation) ([]ReservationResult, error)
		GetBalance(ctx context.Context, filter movement.Filter) (types.Quantity, error)

		CreateStock(ctx context.Context, s stock.Stock) (stock.Stock, error)
		UpdateStock(ctx context.Context, s stock.Stock) error
		RemoveStock(ctx context.Context, id types.ID) error
		ListStock(ctx context.Context) ([]stock.Stock, error)
		GetStockByID(ctx context.Context, id types.ID) (stock.Stock, error)

		ListMovements(ctx context.Context, filter movement.Filter) ([]movement.Movement, error)
		SaveMovement(ctx context.Context, resourceID types.ID, movement movement.Movement) (AdditionResult, error)
		RemoveMovement(ctx context.Context, resourceID types.ID, movementID types.ID) error
	}

	workflow struct {
		stocks    stock.Actions
		movements movement.Actions
	}
)

func NewWorkflow(stocks stock.Actions, movements movement.Actions) Workflow {
	return workflow{
		stocks:    stocks,
		movements: movements,
	}
}
