package start

import (
	"gomies/app/core/workflows/category"
	crew "gomies/app/core/workflows/crew"
	"gomies/app/core/workflows/product"
	"gomies/app/core/wrappers"
)

func NewWorkflows(actions wrappers.Actions, managers wrappers.Managers) wrappers.Workflows {
	products := product.NewWorkflow(actions.Products, actions.Stocks, actions.Categories, managers.Transaction)
	categories := category.NewWorkflow(actions.Categories, managers.Transaction)
	cr := crew.NewWorkflow(actions.Stores, actions.Crew, managers.Transaction, managers.Authorization)

	return wrappers.Workflows{
		Products:   products,
		Categories: categories,
		Crew:       cr,
	}
}
