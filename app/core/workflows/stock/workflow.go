package stock

import (
	"gomies/app/core/entities/stocking/stock"
)

var _ stock.Workflow = workflow{}

func NewWorkflow(stocks stock.Actions) stock.Workflow {
	return workflow{
		stocks: stocks,
	}
}

type workflow struct {
	stocks stock.Actions
}
