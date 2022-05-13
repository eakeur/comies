package category

import (
	"context"
	category2 "gomies/app/core/entities/category"
	"gomies/app/sdk/fault"
)

func (w workflow) GetCategory(ctx context.Context, key category2.Key) (category2.Category, error) {
	const operation = "Workflows.Product.GetCategory"

	ct, err := w.categories.GetCategory(ctx, key)
	if err != nil {
		return category2.Category{}, fault.Wrap(err, operation)
	}

	return ct, nil
}
