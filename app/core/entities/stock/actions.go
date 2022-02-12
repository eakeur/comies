package stock

import (
	"context"
	"gomies/app/core/types/id"
)

type Actions interface {
	ComputeStock(context.Context, Filter) (Actual, error)
	GetMovement(context.Context, id.External) (Movement, error)
	ListMovements(context.Context, Filter) ([]Movement, error)
	AddToStock(context.Context, Movement) (Movement, error)
	RemoveFromStock(context.Context, id.External) error
}
