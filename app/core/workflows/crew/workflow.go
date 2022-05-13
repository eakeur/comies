package crew

import (
	"context"
	crew2 "gomies/app/core/entities/crew"
	"gomies/app/core/entities/store"
	"gomies/app/sdk/session"
)

//go:generate moq -fmt goimports -out workflow_mock.go . Workflow:WorkflowMock

type (
	Workflow interface {
		CreateMember(ctx context.Context, op crew2.Member) (crew2.Member, error)
		ListMembers(ctx context.Context, operatorFilter crew2.Filter) ([]crew2.Member, int, error)
		GetMember(ctx context.Context, key crew2.Key) (crew2.Member, error)
		RemoveMember(ctx context.Context, key crew2.Key) error
		UpdateMember(ctx context.Context, op crew2.Member) error
		AuthenticateMember(ctx context.Context, auth AuthRequest) (session.Session, error)
	}

	AuthRequest struct {
		Nickname       string
		Password       string
		PersistSession bool
	}

	workflow struct {
		stores   store.Actions
		crew     crew2.Actions
		sessions session.Manager
	}
)

var _ Workflow = workflow{}

func NewWorkflow(
	stores store.Actions,
	crew crew2.Actions,
	sessions session.Manager,
) Workflow {
	return workflow{
		stores:   stores,
		crew:     crew,
		sessions: sessions,
	}
}
