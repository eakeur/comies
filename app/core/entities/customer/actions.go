package customer

import (
	"context"
	"gomies/pkg/sdk/types"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type (
	Actions interface {
		ListCustomers(ctx context.Context, f Filter) ([]Customer, error)
		GetCustomers(ctx context.Context, uid types.UID) (Customer, error)
		CreateCustomer(ctx context.Context, c Customer) (Customer, error)
		UpdateCustomer(ctx context.Context, c Customer) error
		RemoveCustomer(ctx context.Context, uid types.UID) error
	}
)
