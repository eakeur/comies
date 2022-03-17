package crew

import (
	"context"
	"gomies/pkg/iam/core/entities/crew"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/session"
)

func (w workflow) Remove(ctx context.Context, key crew.Key) error {
	const operation = "Workflows.Crew.Remove"
	ctx = w.transactions.Begin(ctx)
	defer w.transactions.End(ctx)

	_, err := session.DelegateSessionProps(ctx, operation, &key.Store, nil)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	err = w.crew.Remove(ctx, key)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	return nil
}
