package stock

import (
	"context"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/types"
	"gomies/pkg/stocking/core/entities/stock"
)

func (w workflow) Compute(ctx context.Context, filter stock.Filter) (types.Quantity, error) {
	const operation = "Workflows.Stock.Compute"

	if err := filter.Validate(); err != nil {
		return 0, fault.Wrap(err, operation)
	}

	actual, err := w.stocks.Compute(ctx, filter)
	if err != nil {
		return 0, fault.Wrap(err, operation)
	}

	return actual, nil

}
