package crew

import (
	"context"
	"gomies/app/core/entities/iam/crew"
	"gomies/app/core/entities/iam/store"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/session"
	"strings"
)

func (w workflow) AuthenticateMember(ctx context.Context, auth AuthRequest) (session.Session, error) {
	const operation = "Workflows.Crew.AuthenticateMember"

	operatorNick, storeNick := split(auth.Nickname)
	if operatorNick == "" || storeNick == "" {
		return session.Session{}, crew.ErrInvalidAuthArguments
	}

	op, err := w.crew.GetMemberWithNicknames(ctx, operatorNick, storeNick)
	if err != nil {
		return session.Session{}, fault.Wrap(err, operation)
	}

	if err := op.Password.Compare(auth.Password); err != nil {
		return session.Session{}, err
	}

	pref, err := w.stores.ListPreferences(ctx, store.Key{ID: op.Store.StoreID})
	if err != nil {
		return session.Session{}, fault.Wrap(err, operation)
	}

	_, ses, err := w.sessions.Create(ctx, session.Session{
		OperatorID:   op.ID,
		StoreID:      op.Store.StoreID,
		OperatorName: op.Name,
		Permissions:  op.Permissions,
		Preferences:  pref,
	})
	if err != nil {
		return session.Session{}, fault.Wrap(err, operation)
	}

	return ses, nil
}

func split(mail string) (string, string) {
	arr := strings.Split(mail, "@")
	if len(arr) < 2 {
		return "", ""
	}
	return arr[0], arr[1]
}
