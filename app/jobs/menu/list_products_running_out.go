package menu

import (
	"comies/app/core/menu/movement"
	"comies/app/core/menu/product"
	"comies/app/core/types"
	"context"
	"sync"
)

type MissingProduct struct {
	product.Product
	StockBalance     types.Quantity
	BalancePercetage float64
}

func (w jobs) ListProductsRunningOut(ctx context.Context) ([]MissingProduct, error) {
	list, err := w.products.List(ctx, product.Filter{
		Types: []types.Type{product.InputType, product.OutputType},
	})
	if err != nil {
		return nil, err
	}

	wg := sync.WaitGroup{}
	missing := make([]MissingProduct, len(list))
	for i, p := range list {
		wg.Add(1)

		go func(i int, p product.Product) {
			defer wg.Done()

			bal, err := w.movements.Balance(ctx, movement.Filter{ProductID: p.ID})
			if err != nil || p.MaximumQuantity/4 <= bal {
				return
			}

			missing[i] = MissingProduct{
				Product: p,
			}
		}(i, p)
	}

	wg.Wait()

	return missing, nil

}
