package movement

import (
	"context"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	List(ctx context.Context, filter Filter) ([]Movement, error)
	Create(ctx context.Context, m Movement) error
}
