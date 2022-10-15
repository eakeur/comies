package menu

import (
	"comies/app/core/id"
	"comies/app/core/menu"
	"comies/app/core/types"
	"comies/app/data/ids"
	"comies/app/data/movements"
	"comies/app/data/products"
	"context"
	"time"
)

type ActualBalance struct {
	ID    id.ID
	Count types.Quantity
}

func CreateMovement(ctx context.Context, m menu.Movement) (ActualBalance, error) {

	m.ID = ids.Create()

	if m.Date.IsZero() {
		m.Date = time.Now().UTC()
	}

	if err := menu.ValidateMovement(m); err != nil {
		return ActualBalance{}, err
	}

	prd, err := products.GetByID(ctx, m.ProductID)
	if err != nil {
		return ActualBalance{}, err
	}

	if err := menu.CheckMovementTypeCompatibility(m.Type, prd.Type); err != nil {
		return ActualBalance{}, err
	}

	actual, err := movements.GetBalance(ctx, menu.MovementFilter{})
	if err != nil {
		return ActualBalance{}, err
	}

	actual = menu.IncrementStockQuantity(actual, m)
	if err := menu.CanStockAfford(m, actual, prd.Stock); err != nil {
		return ActualBalance{}, err
	}

	return ActualBalance{
		ID:    m.ID,
		Count: actual,
	}, movements.Create(ctx, m)
}
