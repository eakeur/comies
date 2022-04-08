package customer

import (
	"context"
	"gomies/app/core/entities/customer"
	"gomies/pkg/sdk/fault"
)

func (w workflow) UpdateCustomer(ctx context.Context, op customer.Customer) error {
	const operation = "Workflows.Customers.UpdateCustomer"

	if err := op.Validate(); err != nil {
		return fault.Wrap(err, operation)
	}

	err := w.customers.UpdateCustomer(ctx, op)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	return nil
}
