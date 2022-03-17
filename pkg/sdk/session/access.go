package session

import (
	"context"
	"gomies/pkg/sdk/types"
)

const ContextKey types.ContextKey = "session-context-key"

type Session struct {
	OperatorID      types.External
	StoreExternalID types.External
	StoreInternalID types.ID
	OperatorName    string
	Preferences     types.Preferences
	Permissions     types.Permissions
	Digest          string
}

func (s Session) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, ContextKey, s)
}
