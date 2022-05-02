package customer

import (
	"context"
	"gomies/app/core/entities/customer"
	"gomies/app/sdk/fault"
)

func (w workflow) ListCustomers(ctx context.Context, filter customer.Filter) ([]customer.Customer, int, error) {
	const operation = "Workflows.Customer.ListCustomers"

	ct, count, err := w.customers.ListCustomers(ctx, filter)
	if err != nil {
		return []customer.Customer{}, 0, fault.Wrap(err, operation)
	}

	return ct, count, nil
}
