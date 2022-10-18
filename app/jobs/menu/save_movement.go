package menu

import (
	"comies/app/core/menu"
	"comies/app/core/types"
	"comies/app/data/ids"
	"comies/app/data/movements"
	"comies/app/data/products"
	"context"
	"time"
)

type Balance struct {
	ID    types.ID
	Count types.Quantity
}

func CreateMovement(ctx context.Context, m menu.Movement, q types.Quantity) (Balance, error) {

	m.ID = ids.Create()

	if m.Date.IsZero() {
		m.Date = time.Now().UTC()
	}

	m = menu.AssignMovementQuantity(m, q)

	if err := menu.ValidateMovement(m); err != nil {
		return Balance{}, err
	}

	prd, err := products.GetByID(ctx, m.ProductID)
	if err != nil {
		return Balance{}, err
	}

	if err := menu.CheckMovementTypeCompatibility(m.Type, prd.Type); err != nil {
		return Balance{}, err
	}

	actual, err := movements.GetBalance(ctx, menu.MovementFilter{})
	if err != nil {
		return Balance{}, err
	}

	actual = menu.IncrementStockQuantity(actual, m)
	if err := menu.CanStockAfford(m, actual, prd); err != nil {
		return Balance{}, err
	}

	return Balance{
		ID:    m.ID,
		Count: actual,
	}, movements.Create(ctx, m)
}
