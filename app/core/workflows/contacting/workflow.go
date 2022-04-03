package contacting

import (
	"gomies/app/core/entities/contacting"
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
