package status

import (
	"comies/app/core/types"
	"context"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	Update(ctx context.Context, s Status) error
	ListHistory(ctx context.Context, orderID types.ID) ([]Status, error)
	GetLastUpdate(ctx context.Context, orderID types.ID) (Status, error)
}