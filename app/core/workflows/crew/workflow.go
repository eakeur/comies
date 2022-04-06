package crew

import (
	"context"
	"gomies/app/core/entities/iam/crew"
	"gomies/app/core/entities/iam/store"
	"gomies/pkg/sdk/session"
)

//go:generate moq -fmt goimports -out workflow_mock.go . Workflow:WorkflowMock

type (
	Workflow interface {
		CreateMember(ctx context.Context, op crew.Member) (crew.Member, error)
		ListMembers(ctx context.Context, operatorFilter crew.Filter) ([]crew.Member, error)
		GetMember(ctx context.Context, key crew.Key) (crew.Member, error)
		RemoveMember(ctx context.Context, key crew.Key) error
		UpdateMember(ctx context.Context, op crew.Member) error
		AuthenticateMember(ctx context.Context, auth AuthRequest) (session.Session, error)
	}

	AuthRequest struct {
		Nickname       string
		Password       string
		PersistSession bool
	}

	workflow struct {
		stores   store.Actions
		crew     crew.Actions
		sessions session.Manager
	}
)

var _ Workflow = workflow{}

func NewWorkflow(
	stores store.Actions,
	crew crew.Actions,
	sessions session.Manager,
) Workflow {
	return workflow{
		stores:   stores,
		crew:     crew,
		sessions: sessions,
	}
}
