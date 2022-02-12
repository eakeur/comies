package product

import (
	"context"
	"gomies/app/core/types/id"
)

func (w workflow) RemoveFromStock(ctx context.Context, ext id.External) error {
	ctx = w.transactions.Begin(ctx)
	defer w.transactions.Rollback(ctx)

	err := w.stocks.RemoveFromStock(ctx, ext)
	if err != nil {
		return err
	}
	w.transactions.Commit(ctx)
	return nil
}
