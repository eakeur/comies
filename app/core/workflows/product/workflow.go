package product

import (
	"gomies/app/core/entities/category"
	"gomies/app/core/entities/product"
	"gomies/app/core/entities/stock"
	"gomies/app/core/managers/transaction"
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
