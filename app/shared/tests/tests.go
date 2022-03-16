package tests

import (
	"context"
	"gomies/app/core/managers/session"
	"gomies/app/core/managers/transaction"
	"gomies/app/core/types/id"
	"gomies/app/core/wrappers"
)

func GetFakeManagers() wrappers.Managers {
	return wrappers.Managers{
		Transaction: &transaction.ManagerMock{
			BeginFunc: func(contextMoqParam context.Context) context.Context {
				return contextMoqParam
			},
			CommitFunc:   func(contextMoqParam context.Context) {},
			EndFunc:      func(contextMoqParam context.Context) {},
			RollbackFunc: func(contextMoqParam context.Context) {},
		},
	}
}

func GetAuthorizedContext() context.Context {
	return context.WithValue(context.Background(), session.ContextKey, session.Session{
		OperatorID:      id.ExternalFrom("7a4ad106-f91d-4898-955d-91f0e7e93971"),
		StoreExternalID: id.ExternalFrom("7a4ad106-f91d-4898-955d-91f0e7e93972"),
		StoreInternalID: 1,
		OperatorName:    "Operator Perth",
		Permissions:     "*",
	})
}
