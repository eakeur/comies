package menu

import (
	"comies/app/core/menu/movement"
	"context"
	"time"
)

func (w jobs) CreateMovement(ctx context.Context, m movement.Movement) (movement.Movement, error) {
	save, err := m.
		WithID(w.createID()).WithDate(time.Now()).
		AssertQuantity().Validate()
	if err != nil {
		return movement.Movement{}, err
	}

	p, err := w.products.GetByID(ctx, m.ProductID)
	if err != nil {
		return movement.Movement{}, err
	}

	if (p.IsInput() && m.Type == movement.OutputType) || p.IsComposite() {
		return movement.Movement{}, movement.ErrInvalidProductType
	}

	return save, w.movements.Create(ctx, save)
}
