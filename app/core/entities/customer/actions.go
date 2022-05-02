package customer

import (
	"context"
	"gomies/app/sdk/types"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type (
	Actions interface {
		ListCustomers(ctx context.Context, f Filter) ([]Customer, int, error)
		GetCustomer(ctx context.Context, id types.ID) (Customer, error)
		CreateCustomer(ctx context.Context, c Customer) (Customer, error)
		UpdateCustomer(ctx context.Context, c Customer) error
		RemoveCustomer(ctx context.Context, id types.ID) error
	}
)
