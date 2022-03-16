package crew

import (
	"context"
	"gomies/app/core/managers/session"
	"gomies/app/core/types/id"
)

//go:generate moq -fmt goimports -out workflow_mock.go . Workflow:WorkflowMock
type Workflow interface {
	Create(context.Context, Operator) (Operator, error)
	List(context.Context, Filter) ([]Operator, error)
	Get(context.Context, id.External) (Operator, error)
	Remove(context.Context, id.External) error
	Update(context.Context, Operator) error
	Authenticate(context.Context, AuthRequest) (session.Session, error)
}

type AuthRequest struct {
	Nickname       string
	Password       string
	PersistSession bool
}
