package category

import (
	"gomies/app/core/entities/category"
	"gomies/app/core/managers/transaction"
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
