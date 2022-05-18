package crew

import (
	"context"
	"gomies/app/core/entities/store"
	"gomies/app/sdk/fault"
)

func (w workflow) CreateStore(ctx context.Context, st store.Store) (store.Store, error) {

	st, err := w.stores.CreateStore(ctx, st)
	if err != nil {
		return store.Store{}, fault.Wrap(err)
	}

	return st, nil
}
