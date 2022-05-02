package customer

import (
	"context"
	"gomies/app/core/entities/customer"
	"gomies/app/sdk/fault"
)

func (w workflow) CreateCustomer(ctx context.Context, c customer.Customer) (customer.Customer, error) {
	const operation = "Workflows.Customer.CreateCustomer"

	if err := c.Validate(); err != nil {
		return customer.Customer{}, fault.Wrap(err, operation)
	}

	c, err := w.customers.CreateCustomer(ctx, c)
	if err != nil {
		return customer.Customer{}, fault.Wrap(err, operation)
	}

	return c, nil
}
