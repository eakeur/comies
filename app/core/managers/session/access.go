package session

import (
	"context"
	"gomies/app/core/entities/preferences"
	"gomies/app/core/types/id"
	"gomies/app/core/types/key"
	"gomies/app/core/types/permission"
)

const ContextKey key.ContextKey = "session-context-key"

type Session struct {
	OperatorID      id.External
	StoreExternalID id.External
	StoreInternalID id.ID
	OperatorName    string
	Preferences     preferences.Preferences
	Permissions     permission.Permissions
	Digest          string
}

func (s Session) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, ContextKey, s)
}
