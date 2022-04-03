package crew

import (
	"gomies/app/core/entities/iam/crew"
	"gomies/app/core/entities/iam/store"
	"gomies/pkg/sdk/session"
)

var _ crew.Workflow = workflow{}

func NewWorkflow(
	stores store.Actions,
	crew crew.Actions,
	sessions session.Manager,
) crew.Workflow {
	return workflow{
		stores:   stores,
		crew:     crew,
		sessions: sessions,
	}
}

type workflow struct {
	stores   store.Actions
	crew     crew.Actions
	sessions session.Manager
}
