package stocking

import (
	"comies/app/core/entities/stock"
	"comies/app/sdk/throw"
	"context"
)

func (w workflow) CreateStock(ctx context.Context, s stock.Stock) (stock.Stock, error) {

	w.id.Create(&s.ID)

	s, err := w.stocks.Create(ctx, s)
	if err != nil {
		return stock.Stock{}, throw.Error(err)
	}

	return s, nil
}
