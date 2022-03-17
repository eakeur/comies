package transaction

import "context"

//go:generate moq -fmt goimports -out manager_mock.go . Manager:ManagerMock
type Manager interface {
	// Begin starts a transaction and stores an object to it in this context
	Begin(context.Context) context.Context

	// Commit commits a transaction in this context
	Commit(context.Context)

	// Rollback rollbacks a transaction in this context
	Rollback(context.Context)

	// End finishes a transaction, calling commit if there are no errors or calling rollback if there are any
	End(ctx context.Context)
}
