package customer

import (
	"context"
	"gomies/app/core/entities/customer"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) GetCustomer(ctx context.Context, id types.ID) (customer.Customer, error) {
	const operation = "Workflows.Customer.GetCustomer"

	if id.Empty() {
		return customer.Customer{}, fault.Wrap(fault.ErrMissingUID, operation)
	}

	c, err := w.customers.GetCustomer(ctx, id)
	if err != nil {
		return customer.Customer{}, fault.Wrap(err, operation)
	}

	return c, nil
}
