package stock

import (
	"gomies/pkg/stocking/core/entities/stock"
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
