package menu

import (
	"comies/app/core/entities/movement"
	"comies/app/core/entities/product"
	"comies/app/sdk/throw"
	"context"
)

func (w workflow) CreateMovement(ctx context.Context, mv movement.Movement) (ActualBalance, error) {
	actual, err := w.movements.GetBalanceByProductID(ctx, mv.ProductID, movement.Filter{})
	if err != nil {
		return ActualBalance{}, throw.Error(err)
	}

	stk, err := w.products.GetStockInfoByID(ctx, mv.ProductID)
	if err != nil {
		return ActualBalance{}, throw.Error(err)
	}

	w.id.Create(&mv.ID)

	if err := mv.Validate(); err != nil {
		return ActualBalance{}, throw.Error(err)
	}

	actual += mv.Value()
	if actual > stk.MaximumQuantity {
		return ActualBalance{}, throw.Error(product.ErrStockAlreadyFull)
	}

	if actual < stk.MinimumQuantity {
		return ActualBalance{}, throw.Error(product.ErrStockNegative)
	}

	_, err = w.movements.Create(ctx, mv)
	if err != nil {
		return ActualBalance{}, throw.Error(err)
	}

	return ActualBalance{
		ID:    mv.ID,
		Count: actual,
	}, nil
}
