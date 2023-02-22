package menu

import (
	"comies/core/menu/movement"
	"comies/core/menu/product"
	"comies/core/types"
	"context"

	"golang.org/x/sync/errgroup"
)

type Saleable struct {
	ID    types.ID
	Code  string
	Name  string
	Stock types.Quantity
	Price types.Currency
}

func (w jobs) ListSaleable(ctx context.Context, identifier string) ([]Saleable, error) {

	products, err := w.products.List(ctx, product.Filter{
		Code: identifier,
		Name: identifier,
		Types: []types.Type{
			product.OutputCompositeType,
			product.OutputType,
		},
	})

	if err != nil {
		return nil, err
	}

	eg, gtx := errgroup.WithContext(ctx)
	saleable := make([]Saleable, len(products))

	for i, p := range products {
		i, p := i, p

		eg.Go(func() (err error) {
			var (
				qt types.Quantity
				pr types.Currency
			)

			ieg := errgroup.Group{}

			ieg.Go(func() (err error) {
				pr, err = w.prices.GetLatestValue(gtx, p.ID)
				return
			})

			if !p.IsComposite() {
				ieg.Go(func() (err error) {
					qt, err = w.movements.Balance(gtx, movement.Filter{ProductID: p.ID})
					return
				})
			}

			if err = ieg.Wait(); err != nil {
				return
			}

			saleable[i] = Saleable{
				ID:    p.ID,
				Code:  p.Code,
				Name:  p.Name,
				Stock: qt,
				Price: pr,
			}

			return
		})
	}

	if err = eg.Wait(); err != nil {
		return nil, err
	}

	return saleable, nil

}
