package stock

import (
	"context"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/types"
	"gomies/pkg/stocking/core/entities/stock"
)

func (w workflow) ComputeSome(ctx context.Context, filter stock.Filter, resourcesIDs ...types.UID) ([]types.Quantity, error) {
	const operation = "Workflows.Stock.ComputeSome"

	if err := filter.Validate(); err != nil {
		return []types.Quantity{}, fault.Wrap(err, operation)
	}

	actual, err := w.stocks.ComputeSome(ctx, filter, resourcesIDs...)
	if err != nil {
		return []types.Quantity{}, fault.Wrap(err, operation)
	}

	return actual, nil
}
