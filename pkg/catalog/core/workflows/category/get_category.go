package category

import (
	"context"
	"gomies/pkg/catalog/core/entities/category"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/session"
)

func (w workflow) GetCategory(ctx context.Context, key category.Key) (category.Category, error) {
	const operation = "Workflows.Product.GetCategory"

	_, err := session.DelegateSessionProps(ctx, operation, &key.Store, nil)
	if err != nil {
		return category.Category{}, fault.Wrap(err, operation)
	}

	ct, err := w.categories.Get(ctx, key)
	if err != nil {
		return category.Category{}, fault.Wrap(err, operation)
	}

	return ct, nil
}
