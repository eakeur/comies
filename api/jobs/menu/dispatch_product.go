package menu

import (
	"comies/core/menu/movement"
	"comies/core/types"
	"context"
	"errors"
	"time"

	"golang.org/x/sync/errgroup"
)

type Dispatcher struct {
	ProductID types.ID
	AgentID   types.ID
	Quantity  types.Quantity
	Price     types.Currency
	Date      time.Time
}

func (j jobs) DispatchProduct(ctx context.Context, d Dispatcher) error {
	ingredients, err := j.ingredients.ListByProductID(ctx, d.ProductID)
	if err != nil {
		return err
	}

	if len(ingredients) <= 0 {
		_, err := j.CreateMovement(ctx, movement.Movement{
			ProductID: d.ProductID,
			Type:      movement.OutputType,
			AgentID:   d.AgentID,
			Quantity:  d.Quantity,
			Date:      d.Date,
		})
		if err != nil && !errors.Is(err, movement.ErrInvalidProductType) {
			return err
		}
	}

	eg, ctx := errgroup.WithContext(ctx)

	for _, i := range ingredients {
		i := i
		eg.Go(func() error {
			return j.DispatchProduct(ctx, Dispatcher{
				ProductID: i.IngredientID,
				AgentID:   d.AgentID,
				Quantity:  d.Quantity * i.Quantity,
				Price:     d.Price,
				Date:      d.Date,
			})
		})
	}

	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}
