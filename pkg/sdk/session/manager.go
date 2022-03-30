package session

import (
	"context"
)

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
