package contacting

import (
	"context"
	"gomies/pkg/sdk/types"
)

type Workflow interface {
	SaveContact(ctx context.Context, contact Contact) (Contact, error)
	GetAddress(ctx context.Context, targetID types.UID, id types.UID) (Address, error)
	GetPhone(ctx context.Context, targetID types.UID, id types.UID) (Phone, error)
	ListAddresses(ctx context.Context, targetID types.UID) ([]Address, error)
	ListPhones(ctx context.Context, targetID types.UID) ([]Phone, error)
	RemovePhone(ctx context.Context, targetID types.UID, id types.UID) error
	RemoveAddress(ctx context.Context, targetID types.UID, id types.UID) error
}
