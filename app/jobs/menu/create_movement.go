package menu

import (
	"comies/app/core/menu/movement"
	"comies/app/core/types"
	"context"
	"time"
)

func (w jobs) CreateMovement(ctx context.Context, m movement.Movement) (types.ID, error) {
	save, err := m.
		WithID(w.createID()).WithDate(time.Now()).
		AssertQuantity().Validate()
	if err != nil {
		return 0, err
	}

	p, err := w.products.GetByID(ctx, m.ProductID)
	if err != nil {
		return 0, err
	}

	if (p.IsInput() && m.Type == movement.OutputType) || p.IsComposite() {
		return 0, movement.ErrInvalidProductType
	}

	return save.ID, w.movements.Create(ctx, save)
}
