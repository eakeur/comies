package contacting

import (
	"context"
	"gomies/app/core/entities/address"
	"gomies/app/core/entities/phone"
	"gomies/app/sdk/types"
)

var _ Workflow = workflow{}

func NewWorkflow(addresses address.Actions, phones phone.Actions) Workflow {
	return workflow{
		phones:    phones,
		addresses: addresses,
	}
}

type (
	Workflow interface {
		CreateAddresses(ctx context.Context, addresses []address.Address) ([]address.Address, error)
		CreatePhones(ctx context.Context, phones []phone.Phone) ([]phone.Phone, error)
		GetAddressByID(ctx context.Context, id types.ID) (address.Address, error)
		GetPhoneByID(ctx context.Context, id types.ID) (phone.Phone, error)
		ListAddresses(ctx context.Context, targetID types.ID) ([]address.Address, error)
		ListPhones(ctx context.Context, targetID types.ID) ([]phone.Phone, error)
		RemovePhone(ctx context.Context, phoneID types.ID) error
		RemoveAddress(ctx context.Context, addressID types.ID) error
		RemoveTargetsPhones(ctx context.Context, targetID types.ID) error
		RemoveTargetsAddresses(ctx context.Context, targetID types.ID) error
	}

	workflow struct {
		phones    phone.Actions
		addresses address.Actions
	}
)
