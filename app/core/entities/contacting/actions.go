package contacting

import (
	"context"
	"gomies/app/sdk/types"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	ListAddresses(ctx context.Context, target types.ID) ([]Address, error)
	GetAddress(ctx context.Context, target types.ID, addressID types.ID) (Address, error)
	SaveAddresses(ctx context.Context, target types.ID, addresses ...Address) ([]Address, error)
	RemoveAddresses(ctx context.Context, target types.ID, ids ...types.ID) error
	ListPhones(ctx context.Context, target types.ID) ([]Phone, error)
	GetPhone(ctx context.Context, target types.ID, id types.ID) (Phone, error)
	SavePhones(ctx context.Context, target types.ID, phones ...Phone) ([]Phone, error)
	RemovePhones(ctx context.Context, target types.ID, ids ...types.ID) error
}
