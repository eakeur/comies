package contacting

import (
	"context"
	"gomies/pkg/sdk/types"
)

type Workflow interface {
	SaveContact(ctx context.Context, contact Contact) (Contact, error)
	GetAddress(ctx context.Context, targetID types.UID, id types.UID) (Address, error)
	GetPhone(ctx context.Context, targetID types.UID, id types.UID) (Phone, error)
	GetContact(ctx context.Context, targetID types.UID) (Contact, error)
	RemovePhones(ctx context.Context, targetID types.UID, id ...types.UID) error
	RemoveAddresses(ctx context.Context, targetID types.UID, id ...types.UID) error
}
