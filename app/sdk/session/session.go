package session

import (
	"context"
	"gomies/app/sdk/types"
)

const ContextKey types.ContextKey = "session-context-key"

type Session struct {
	OperatorID   types.ID
	StoreID      types.ID
	OperatorName string
	Preferences  types.Preferences
	Permissions  types.Permissions
	Digest       string
}

func (s Session) Delegate(store *types.Store, history *types.History) {
	if store != nil && store.StoreID.Empty() {
		store.StoreID = s.StoreID
	}

	if history != nil {
		history.By = s.OperatorID
	}
}

func (s Session) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, ContextKey, s)
}

// FromContext fetches a session from the context and verifies if it is allowed to execute
// the given operation
//
// Possible errors:
//   - ErrNoSession: if there is no session in this context
func FromContext(ctx context.Context) (Session, error) {
	session, ok := ctx.Value(ContextKey).(Session)
	if !ok {
		return Session{}, ErrNoSession
	}

	return session, nil
}

func DelegateFromContext(ctx context.Context, store *types.Store, history *types.History) (Session, error) {
	s, err := FromContext(ctx)
	if err != nil {
		return Session{}, err
	}

	s.Delegate(store, history)

	return s, nil
}
