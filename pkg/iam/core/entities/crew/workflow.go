package crew

import (
	"context"
	"gomies/pkg/sdk/session"
	"gomies/pkg/sdk/types"
)

//go:generate moq -fmt goimports -out workflow_mock.go . Workflow:WorkflowMock

type (
	Workflow interface {
		// Save creates a new operator or updates an existing one
		//
		// Possible errors
		//   - session.ErrNoSession: if there is no session in this context
		//   - permission.ErrNotAllowed: if the session owner is not allowed to perform this operation
		//   - fault.ErrAlreadyExists: if the operator already exists and the "overwrite" flag was not set
		Save(ctx context.Context, op Member, flag ...types.WritingFlag) (Member, error)

		// List searches all operators with the given filter
		//
		// Possible errors:
		//   - session.ErrNoSession: if there is no session in this context
		//   - permission.ErrNotAllowed: if the session owner is not allowed to perform this operation
		List(ctx context.Context, operatorFilter Filter) ([]Member, error)

		// Get retrieves an operator with this key
		//
		// Possible errors:
		//   - session.ErrNoSession: if there is no session in this context
		//   - permission.ErrNotAllowed: if the session owner is not allowed to perform this operation
		//   - fault.ErrNotFound: if the operator does not exist
		Get(ctx context.Context, key Key) (Member, error)

		// Remove deletes an operator
		//
		// Possible errors:
		//   - session.ErrNoSession: if there is no session in this context
		//   - permission.ErrNotAllowed: if the session owner is not allowed to perform this operation
		Remove(ctx context.Context, key Key) error

		// Authenticate creates an access instance to the operator referred to by the nickname and password
		//
		// Possible errors:
		//   - fault.ErrNotFound: if the operator does not exist
		//   - ErrWrongPassword: if the password is not right
		Authenticate(ctx context.Context, auth AuthRequest) (session.Session, error)
	}

	AuthRequest struct {
		Nickname       string
		Password       string
		PersistSession bool
	}
)
