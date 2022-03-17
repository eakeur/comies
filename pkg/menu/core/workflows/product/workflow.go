package product

import (
	"gomies/pkg/menu/core/entities/category"
	"gomies/pkg/menu/core/entities/product"
	"gomies/pkg/sdk/transaction"
	"gomies/pkg/stocking/core/entities/stock"
)

var _ product.Workflow = workflow{}

func NewWorkflow(
	products product.Actions,
	stocks stock.Actions,
	categories category.Actions,
	transaction transaction.Manager,
) product.Workflow {
	return workflow{
		products:     products,
		stocks:       stocks,
		categories:   categories,
		transactions: transaction,
	}
}

type workflow struct {
	products     product.Actions
	stocks       stock.Actions
	categories   category.Actions
	transactions transaction.Manager
}
