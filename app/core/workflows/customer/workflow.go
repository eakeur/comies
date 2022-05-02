package customer

import (
	"context"
	"gomies/app/core/entities/customer"
	"gomies/app/sdk/types"
)

var _ Workflow = workflow{}

type (
	Workflow interface {
		CreateCustomer(ctx context.Context, op customer.Customer) (customer.Customer, error)
		ListCustomers(ctx context.Context, operatorFilter customer.Filter) ([]customer.Customer, int, error)
		GetCustomer(ctx context.Context, id types.ID) (customer.Customer, error)
		RemoveCustomer(ctx context.Context, id types.ID) error
		UpdateCustomer(ctx context.Context, op customer.Customer) error
	}

	workflow struct {
		customers customer.Actions
	}
)

func NewWorkflow(customers customer.Actions) Workflow {
	return workflow{
		customers: customers,
	}
}
