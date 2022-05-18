package content

import (
	"context"
	"gomies/app/sdk/types"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	Create(ctx context.Context, c Content) (Content, error)
	List(ctx context.Context, itemID types.ID) ([]Content, error)
}
