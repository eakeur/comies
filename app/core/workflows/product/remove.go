package product

import (
	"context"
	"gomies/app/core/types/id"
)

func (w workflow) Remove(ctx context.Context, ext id.External) error {
	const operation = "Workflows.Product.Remove"
	w.logger.Debug(ctx, operation, "starting process")
	ctx = w.transactions.Begin(ctx)
	defer w.transactions.Rollback(ctx)

	err := w.products.Remove(ctx, ext)
	if err != nil {
		w.logger.Warn(ctx, operation, err.Error())
		return err
	}

	w.transactions.Commit(ctx)
	w.logger.Debug(ctx, operation, "finished process")
	return err

}
