package stock

import (
	"context"
	"gomies/app/core/types/id"
)

var _ Actions = ActionsMock{}

type ActionsMock struct {
	ComputeStockResult  Actual
	GetMovementResult   Movement
	ListMovementsResult []Movement
	AddToStockResult    Movement
	Error               error
}

func (a ActionsMock) ComputeStock(_ context.Context, _ Filter) (Actual, error) {
	return a.ComputeStockResult, a.Error
}

func (a ActionsMock) GetMovement(_ context.Context, _ id.External) (Movement, error) {
	return a.GetMovementResult, a.Error
}

func (a ActionsMock) ListMovements(_ context.Context, _ Filter) ([]Movement, error) {
	return a.ListMovementsResult, a.Error
}

func (a ActionsMock) AddToStock(_ context.Context, _ Movement) (Movement, error) {
	return a.AddToStockResult, a.Error
}

func (a ActionsMock) RemoveFromStock(_ context.Context, _ id.External) error {
	return a.Error
}
