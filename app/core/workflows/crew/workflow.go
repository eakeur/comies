package crew

import (
	"context"
	"gomies/app/core/entities/crew"
	"gomies/app/core/entities/store"
	"gomies/app/core/managers/session"
	"gomies/app/core/managers/transaction"
	"gomies/app/core/types/id"
)

var _ crew.Workflow = workflow{}

func NewWorkflow(
	stores store.Actions,
	crew crew.Actions,
	transactions transaction.Manager,
	sessions session.Manager,
) crew.Workflow {
	return workflow{
		stores:       stores,
		crew:         crew,
		transactions: transactions,
		sessions:     sessions,
	}
}

type workflow struct {
	stores       store.Actions
	crew         crew.Actions
	transactions transaction.Manager
	sessions     session.Manager
}

func (w workflow) Create(ctx context.Context, operator crew.Operator) (crew.Operator, error) {
	//TODO implement me
	panic("implement me")
}

func (w workflow) List(ctx context.Context, filter crew.Filter) ([]crew.Operator, error) {
	//TODO implement me
	panic("implement me")
}

func (w workflow) Get(ctx context.Context, external id.External) (crew.Operator, error) {
	//TODO implement me
	panic("implement me")
}

func (w workflow) Remove(ctx context.Context, external id.External) error {
	//TODO implement me
	panic("implement me")
}

func (w workflow) Update(ctx context.Context, operator crew.Operator) error {
	//TODO implement me
	panic("implement me")
}
