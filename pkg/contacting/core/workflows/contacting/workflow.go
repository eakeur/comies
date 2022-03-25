package contacting

import (
	"context"
	"gomies/pkg/contacting/core/entities/contacting"
	"gomies/pkg/sdk/types"
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

func (w workflow) GetAddress(ctx context.Context, targetID types.UID, id types.UID) (contacting.Address, error) {
	//TODO implement me
	panic("implement me")
}

func (w workflow) GetPhone(ctx context.Context, targetID types.UID, id types.UID) (contacting.Phone, error) {
	//TODO implement me
	panic("implement me")
}

func (w workflow) GetContact(ctx context.Context, targetID types.UID) (contacting.Contact, error) {
	//TODO implement me
	panic("implement me")
}

func (w workflow) RemovePhones(ctx context.Context, targetID types.UID, id ...types.UID) error {
	//TODO implement me
	panic("implement me")
}

func (w workflow) RemoveAddresses(ctx context.Context, targetID types.UID, id ...types.UID) error {
	//TODO implement me
	panic("implement me")
}
