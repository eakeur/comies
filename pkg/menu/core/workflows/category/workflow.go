package category

import (
	"gomies/pkg/menu/core/entities/category"
	"gomies/pkg/sdk/transaction"
)

var _ category.Workflow = workflow{}

func NewWorkflow(categories category.Actions, transactions transaction.Manager) category.Workflow {
	return workflow{
		categories:   categories,
		transactions: transactions,
	}
}

type workflow struct {
	categories   category.Actions
	transactions transaction.Manager
}
