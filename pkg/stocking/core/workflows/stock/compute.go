package stock

import (
	"context"
	"errors"
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

func (w workflow) ComputeSome(ctx context.Context, filter stock.Filter, resourcesIDs ...types.UID) ([]types.Quantity, error) {
	const operation = "Workflows.Stock.ComputeSome"

	if err := filter.Validate(); err != nil && !errors.Is(err, stock.ErrMissingResourceID) {
		return []types.Quantity{}, fault.Wrap(err, operation)
	}

	for _, id := range resourcesIDs {
		if id.Empty() {
			return []types.Quantity{}, fault.Wrap(stock.ErrMissingResourceID, operation)
		}
	}

	actual, err := w.stocks.ComputeSome(ctx, filter, resourcesIDs...)
	if err != nil {
		return []types.Quantity{}, fault.Wrap(err, operation)
	}

	return actual, nil
}
