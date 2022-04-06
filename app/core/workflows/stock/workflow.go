package stock

import (
	"context"
	"gomies/app/core/entities/stocking/stock"
	"gomies/pkg/sdk/types"
)

var _ Workflow = workflow{}

//go:generate moq -fmt goimports -out workflow_mock.go . Workflow:WorkflowMock
type (
	Workflow interface {
		Compute(ctx context.Context, filter stock.Filter) (types.Quantity, error)
		ComputeSome(ctx context.Context, filter stock.Filter, resourcesIDs ...types.UID) ([]types.Quantity, error)
		ListMovements(ctx context.Context, filter stock.Filter) ([]stock.Movement, error)
		SaveMovements(ctx context.Context, config stock.Config, resourceID types.UID, movements ...stock.Movement) (stock.AdditionResult, error)
		RemoveMovement(ctx context.Context, resourceID types.UID, movementID types.UID) error
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
