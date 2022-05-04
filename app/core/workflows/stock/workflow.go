package stock

import (
	"context"
	"gomies/app/core/entities/stock"
	"gomies/app/sdk/listing"
	"gomies/app/sdk/types"
)

var _ Workflow = workflow{}

//go:generate moq -fmt goimports -out workflow_mock.go . Workflow:WorkflowMock
type (
	Workflow interface {
		ConsumeResources(ctx context.Context, reservationID types.ID) error
		FreeResources(ctx context.Context, reservationID types.ID) error
		ReserveResources(ctx context.Context, reservationID types.ID, reservations []Reservation) ([]ReservationResult, error)
		ComputeStock(ctx context.Context, filter stock.Filter) (types.Quantity, error)

		CreateStock(ctx context.Context, s stock.Stock) (stock.Stock, error)
		UpdateStock(ctx context.Context, s stock.Stock) error
		RemoveStock(ctx context.Context, id types.ID) error
		ListStock(ctx context.Context, filter listing.Filter) ([]stock.Stock, int, error)
		GetStockByID(ctx context.Context, id types.ID) (stock.Stock, error)

		ListMovements(ctx context.Context, filter stock.Filter) ([]stock.Movement, int, error)
		SaveMovements(ctx context.Context, resourceID types.ID, movement stock.Movement) (stock.AdditionResult, error)
		RemoveMovement(ctx context.Context, resourceID types.ID, movementID types.ID) error
		ClosePeriod(ctx context.Context, filter stock.Filter) error
	}

	workflow struct {
		stocks stock.Actions
	}
)

func NewWorkflow(stocks stock.Actions) Workflow {
	return workflow{
		stocks: stocks,
	}
}
