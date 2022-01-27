package stock

import (
	"context"
	"gomies/app/core/types/id"
)

type Actions interface {
	ComputeStock(context.Context, Filter) (Actual, error)
	GetMovements(context.Context, id.External) (Actual, error)
	ListMovements(context.Context, Filter) ([]Movement, error)
	AddToStock(context.Context, Movement) (Movement, error)
	RemoveFromStock(context.Context, id.External) error
}
