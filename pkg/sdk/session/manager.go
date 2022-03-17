package session

import (
	"context"
	"gomies/pkg/sdk/types"
)

// FromContext fetches a session from the context and verifies if it is allowed to execute
// the given operation
//
// Possible errors:
//   - ErrNoSession: if there is no session in this context
//   - preference.ErrNotAllowed: if the session owner is not allowed to perform this operation
func FromContext(ctx context.Context, operation ...string) (Session, error) {
	session, ok := ctx.Value(ContextKey).(Session)
	if !ok {
		return Session{}, ErrNoSession
	}

	if len(operation) > 0 {
		err := session.Permissions.CheckAllowance(operation[0])
		if err != nil {
			return Session{}, err
		}
	}
	return session, nil
}

// DelegateSessionProps fetches a session from the context, verifies if it is allowed to execute
// the given operation and if so, sets entity's StoreID, StoreExternalID, By, Operation and Active properties
// the same as session's
//
// Possible errors:
//   - ErrNoSession: if there is no session in this context
//   - preference.ErrNotAllowed: if the session owner is not allowed to perform this operation
func DelegateSessionProps(ctx context.Context, operation string, entity *types.Entity) (Session, error) {
	session, err := FromContext(ctx)
	if err != nil {
		return Session{}, err
	}

	if err := session.Permissions.CheckAllowance(operation); err != nil {
		return Session{}, err
	}

	// if entity.StoreID == 0 && (entity.StoreExternalID == types.Nil || (entity.StoreExternalID == types.External{})) {
	// 	entity.StoreExternalID = session.StoreExternalID
	// 	entity.StoreID = session.StoreInternalID
	// 	entity.Active = true
	// }

	entity.History = types.History{
		By:        session.OperatorID,
		Operation: operation,
	}

	return session, nil
}

//go:generate moq -fmt goimports -out manager_mock.go . Manager:ManagerMock
type Manager interface {

	// Create creates a new session based on an operator login and assigns it
	// to the context.
	Create(ctx context.Context, op Session) (context.Context, Session, error)

	// Retrieve fetches a session from the string digest and assigns it
	// to the context.
	//
	// Possible errors:
	//   - ErrSessionInvalidOrExpired: if an error occurs when parsing the digest or the expiration time is over
	Retrieve(ctx context.Context, digest string, updateExpiration bool) (context.Context, Session, error)
}
