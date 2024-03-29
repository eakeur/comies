package status

import (
	"comies/core/types"
	"context"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	Update(ctx context.Context, s Status) error
	ListHistory(ctx context.Context, orderID types.ID) ([]Status, error)
	GetLatestUpdate(ctx context.Context, orderID types.ID) (Status, error)
	CountByStatus(ctx context.Context, s types.Status) (types.Quantity, error)
}
