package contacting

import (
	"context"
	"gomies/app/core/entities/contacting"
	"gomies/app/sdk/types"
)

var _ Workflow = workflow{}

func NewWorkflow(contacts contacting.Actions) Workflow {
	return workflow{
		contacts: contacts,
	}
}

type (
	Workflow interface {
		SaveContact(ctx context.Context, contact contacting.Contact) (contacting.Contact, error)
		GetAddress(ctx context.Context, targetID types.ID, id types.ID) (contacting.Address, error)
		GetPhone(ctx context.Context, targetID types.ID, id types.ID) (contacting.Phone, error)
		ListAddresses(ctx context.Context, targetID types.ID) ([]contacting.Address, error)
		ListPhones(ctx context.Context, targetID types.ID) ([]contacting.Phone, error)
		RemovePhone(ctx context.Context, targetID types.ID, id types.ID) error
		RemoveAddress(ctx context.Context, targetID types.ID, id types.ID) error
	}

	workflow struct {
		contacts contacting.Actions
	}
)
