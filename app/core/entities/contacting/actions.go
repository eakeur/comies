package contacting

import (
	"context"
	"gomies/pkg/sdk/types"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	ListAddresses(ctx context.Context, target types.UID) ([]Address, error)
	GetAddress(ctx context.Context, target types.UID, addressID types.UID) (Address, error)
	SaveAddresses(ctx context.Context, target types.UID, addresses ...Address) ([]Address, error)
	RemoveAddresses(ctx context.Context, target types.UID, ids ...types.UID) error
	ListPhones(ctx context.Context, target types.UID) ([]Phone, error)
	GetPhone(ctx context.Context, target types.UID, id types.UID) (Phone, error)
	SavePhones(ctx context.Context, target types.UID, phones ...Phone) ([]Phone, error)
	RemovePhones(ctx context.Context, target types.UID, ids ...types.UID) error
}
