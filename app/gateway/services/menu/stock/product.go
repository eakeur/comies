package stock

import (
	"comies/app/core/workflows/menu"
	"comies/app/core/workflows/stocking"
)

var _ menu.StockService = service{}

type service struct {
	stocks stocking.Workflow
}

func NewService(stocks stocking.Workflow) menu.StockService {
	return service{stocks: stocks}
}
