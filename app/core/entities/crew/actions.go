package crew

import (
	"context"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	GetMember(ctx context.Context, key Key) (Member, error)
	GetMemberWithNicknames(ctx context.Context, operatorNickname string, storeNickname string) (Member, error)
	ListMembers(ctx context.Context, operatorFilter Filter) ([]Member, int, error)
	CreateMember(ctx context.Context, op Member) (Member, error)
	RemoveMember(ctx context.Context, key Key) error
	UpdateMember(ctx context.Context, op Member) error
}
