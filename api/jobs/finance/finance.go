package finance

import (
	"comies/core/finance/movement"
	"comies/core/types"
	"context"
	"time"
)

//go:generate moq -fmt goimports -out workflow_mock.go . Jobs:WorkflowMock
type Jobs interface {
	CreateFinancialMovement(ctx context.Context, m movement.Movement) (types.ID, error)
	ListFinancialMovement(ctx context.Context, filter movement.Filter) ([]movement.Movement, error)
}

type jobs struct {
	movements movement.Actions
	createID  types.CreateID
}

var _ Jobs = jobs{}

func NewJobs(
	movements movement.Actions,
	createID types.CreateID,
) Jobs {
	return jobs{
		movements: movements,
		createID:  createID,
	}
}

func (j jobs) CreateFinancialMovement(ctx context.Context, m movement.Movement) (types.ID, error) {
	save, err := m.
		WithID(j.createID()).WithDate(time.Now()).Validate()
	if err != nil {
		return 0, err
	}

	return save.ID, j.movements.Create(ctx, save)
}

func (j jobs) ListFinancialMovement(ctx context.Context, filter movement.Filter) ([]movement.Movement, error) {
	if err := filter.Validate(); err != nil {
		return nil, err
	}

	return j.movements.List(ctx, filter)
}
