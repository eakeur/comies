package menu

import (
	"comies/app/core/menu/movement"
	"context"
)

func (w jobs) ListMovements(ctx context.Context, filter movement.Filter) ([]movement.Movement, error) {
	if err := filter.Validate(); err != nil {
		return nil, err
	}

	return w.movements.List(ctx, filter)
}
