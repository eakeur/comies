package category

import (
	"context"
	"gomies/app/core/entities/catalog/category"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/types"
)

func (w workflow) SaveCategory(ctx context.Context, ct category.Category, flag ...types.WritingFlag) (category.Category, error) {
	const operation = "Workflows.Category.SaveCategory"

	if err := ct.Validate(); err != nil {
		return category.Category{}, fault.Wrap(err, operation)
	}

	ct, err := w.categories.Save(ctx, ct, flag...)
	if err != nil {
		return category.Category{}, fault.Wrap(err, operation)
	}

	return ct, nil
}
