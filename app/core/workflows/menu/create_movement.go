package menu

import (
	"comies/app/core/id"
	"comies/app/core/movement"
	"comies/app/core/product"
	"comies/app/data/movements"
	"comies/app/data/products"
	"context"
	"time"
)

func CreateMovement(ctx context.Context, mv movement.Movement) (ActualBalance, error) {
	prd, err := products.GetByID(ctx, mv.ProductID)
	if err != nil {
		return ActualBalance{}, err
	}

	if (prd.Type == product.InputType && mv.Type == movement.OutputType) ||
		(prd.Type == product.OutputCompositeType) || (prd.Type == product.InputCompositeType) {
		return ActualBalance{}, movement.ErrInvalidProductType
	}

	actual, err := movements.GetBalanceByProductID(ctx, mv.ProductID, movement.Filter{})
	if err != nil {
		return ActualBalance{}, err
	}

	id.Create(&mv.ID)

	if mv.Date.IsZero() {
		mv.Date = time.Now().UTC()
	}

	if err := mv.Validate(); err != nil {
		return ActualBalance{}, err
	}

	actual += mv.Value()
	if mv.Type == movement.InputType && actual > prd.MaximumQuantity {
		return ActualBalance{}, product.ErrStockAlreadyFull
	}

	if mv.Type == movement.OutputType && actual < prd.MinimumQuantity {
		return ActualBalance{}, product.ErrStockNegative
	}

	_, err = movements.Create(ctx, mv)
	if err != nil {
		return ActualBalance{}, err
	}

	return ActualBalance{
		ID:    mv.ID,
		Count: actual,
	}, nil
}
