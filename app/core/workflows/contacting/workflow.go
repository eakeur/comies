package contacting

import (
	"context"
	"gomies/app/core/entities/contacting"
	"gomies/pkg/sdk/types"
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
		GetAddress(ctx context.Context, targetID types.UID, id types.UID) (contacting.Address, error)
		GetPhone(ctx context.Context, targetID types.UID, id types.UID) (contacting.Phone, error)
		ListAddresses(ctx context.Context, targetID types.UID) ([]contacting.Address, error)
		ListPhones(ctx context.Context, targetID types.UID) ([]contacting.Phone, error)
		RemovePhone(ctx context.Context, targetID types.UID, id types.UID) error
		RemoveAddress(ctx context.Context, targetID types.UID, id types.UID) error
	}

	workflow struct {
		contacts contacting.Actions
	}
)
