package crew

import (
	"context"
	"gomies/app/core/entities/member"
	"gomies/app/sdk/session"
)

//go:generate moq -fmt goimports -out workflow_mock.go . Workflow:WorkflowMock

type (
	Workflow interface {
		Create(ctx context.Context, op member.Member) (member.Member, error)
		List(ctx context.Context, operatorFilter member.Filter) ([]member.Member, int, error)
		GetByKey(ctx context.Context, key member.Key) (member.Member, error)
		Remove(ctx context.Context, key member.Key) error
		Update(ctx context.Context, op member.Member) error
		AuthenticateMember(ctx context.Context, auth AuthRequest) (session.Session, error)
	}

	AuthRequest struct {
		Nickname       string
		Password       string
		PersistSession bool
	}

	workflow struct {
		crew     member.Actions
		sessions session.Manager
	}
)

var _ Workflow = workflow{}

func NewWorkflow(
	crew member.Actions,
	sessions session.Manager,
) Workflow {
	return workflow{
		crew:     crew,
		sessions: sessions,
	}
}
