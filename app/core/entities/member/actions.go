package member

import (
	"context"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	GetByKey(ctx context.Context, key Key) (Member, error)
	List(ctx context.Context, operatorFilter Filter) ([]Member, error)
	Create(ctx context.Context, op Member) (Member, error)
	Remove(ctx context.Context, key Key) error
}
