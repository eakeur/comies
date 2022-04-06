package store

import (
	"context"
	"gomies/pkg/sdk/types"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	GetStore(ctx context.Context, storeID Key) (Store, error)
	ListStore(ctx context.Context, storeFilter Filter) ([]Store, error)
	CreateStore(ctx context.Context, st Store) (Store, error)
	RemoveStore(ctx context.Context, key Key) error
	UpdateStore(ctx context.Context, st Store) error
	ListPreferences(ctx context.Context, storeKey Key, modules ...string) (types.Preferences, error)
	SavePreferences(ctx context.Context, storeKey Key, pref types.Preferences) (types.Preferences, error)
}
