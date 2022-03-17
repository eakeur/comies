package crew

import (
	"gomies/pkg/iam/core/entities/crew"
	"gomies/pkg/iam/core/entities/store"
	"gomies/pkg/sdk/session"
	"gomies/pkg/sdk/transaction"
)

var _ crew.Workflow = workflow{}

func NewWorkflow(
	stores store.Actions,
	crew crew.Actions,
	transactions transaction.Manager,
	sessions session.Manager,
) crew.Workflow {
	return workflow{
		stores:       stores,
		crew:         crew,
		transactions: transactions,
		sessions:     sessions,
	}
}

type workflow struct {
	stores       store.Actions
	crew         crew.Actions
	transactions transaction.Manager
	sessions     session.Manager
}
