package stock

import (
	"context"
	"gomies/app/core/entities/stock"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) ComputeStock(ctx context.Context, filter stock.Filter) (types.Quantity, error) {
	const operation = "Workflows.Stock.ComputeStock"

	if err := filter.Validate(); err != nil {
		return 0, fault.Wrap(err, operation)
	}

	actual, err := w.stocks.ComputeStock(ctx, filter)
	if err != nil {
		return 0, fault.Wrap(err, operation)
	}

	return actual, nil

}
