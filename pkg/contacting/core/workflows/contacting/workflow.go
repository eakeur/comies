package contacting

import (
	"gomies/pkg/contacting/core/entities/contacting"
)

var _ contacting.Workflow = workflow{}

func NewWorkflow(contacts contacting.Actions) contacting.Workflow {
	return workflow{
		contacts: contacts,
	}
}

type workflow struct {
	contacts contacting.Actions
}
