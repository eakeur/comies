package tests

import (
	"context"
	"gomies/pkg/sdk/session"
	"gomies/pkg/sdk/transaction"
	"gomies/pkg/sdk/types"
)

type (
	ManagersMocks struct {
		Sessions     session.Manager
		Transactions transaction.Manager
	}
)

func WorkflowContext(operatorID, storeID types.UID) context.Context {
	return context.WithValue(
		context.Background(),
		session.ContextKey,
		session.Session{
			OperatorID:   operatorID,
			StoreID:      storeID,
			OperatorName: "Tester Smith",
			Permissions:  "*",
			Digest:       "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Preferences:  types.Preferences{},
		},
	)
}

func Managers() ManagersMocks {
	return ManagersMocks{
		Transactions: &transaction.ManagerMock{
			BeginFunc:    func(contextMoqParam context.Context) context.Context { return contextMoqParam },
			CommitFunc:   func(contextMoqParam context.Context) {},
			EndFunc:      func(contextMoqParam context.Context) {},
			RollbackFunc: func(contextMoqParam context.Context) {},
		},
		Sessions: &session.ManagerMock{
			CreateFunc: func(ctx context.Context, op session.Session) (context.Context, session.Session, error) {
				return ctx, op, nil
			},
			RetrieveFunc: func(ctx context.Context, digest string, updateExpiration bool) (context.Context, session.Session, error) {
				return ctx, session.Session{
					OperatorID:   types.UIDFrom("0b3d0d55-4610-4516-96ec-8667b519599d"),
					StoreID:      types.UIDFrom("0b3d0d55-4610-4516-96ec-8667b519599a"),
					OperatorName: "Tester Smith",
					Permissions:  "*",
					Digest:       "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
					Preferences:  types.Preferences{},
				}, nil
			},
		},
	}
}
