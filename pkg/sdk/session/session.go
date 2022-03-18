package session

import (
	"context"
	"gomies/pkg/sdk/types"
	"strings"
)

const ContextKey types.ContextKey = "session-context-key"

type Session struct {
	OperatorID   types.External
	StoreID      types.External
	OperatorName string
	Preferences  types.Preferences
	Permissions  types.Permissions
	Digest       string
}

func (s Session) Delegate(operation string, entity *types.Store, history *types.History)  {
	if entity != nil && entity.StoreID.Empty() {
		entity.StoreID = s.StoreID
	}

	if history != nil {
		history.By = s.OperatorID
		history.Operation = operation
	}
}

func (s Session) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, ContextKey, s)
}

func (s Session) isAllowed(operation string) bool {
	if s.Permissions == "*" {
		return true
	}

	if !strings.Contains(string(s.Permissions), operation) {
		return false
	}

	return true
}
