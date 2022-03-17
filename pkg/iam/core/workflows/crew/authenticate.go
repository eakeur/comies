package crew

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"gomies/pkg/iam/core/entities/crew"
	"gomies/pkg/iam/core/entities/store"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/session"
	"strings"
)

func (w workflow) Authenticate(ctx context.Context, auth crew.AuthRequest) (session.Session, error) {
	const operation = "Workflows.Crew.Authenticate"

	operatorNick, storeNick := split(auth.Nickname)

	op, err := w.crew.GetWithNicknames(ctx, operatorNick, storeNick)
	if err != nil {
		return session.Session{}, fault.Wrap(err, operation)
	}

	err = bcrypt.CompareHashAndPassword([]byte(op.Password), []byte(auth.Password))
	if err != nil {
		return session.Session{}, crew.ErrWrongPassword
	}

	pref, err := w.stores.ListPreferences(ctx, store.Key{ID: op.StoreExternalID})
	if err != nil {
		return session.Session{}, fault.Wrap(err, operation)
	}

	ctx, ses, err := w.sessions.Create(ctx, session.Session{
		OperatorID:      op.ExternalID,
		StoreExternalID: op.StoreExternalID,
		StoreInternalID: op.StoreID,
		OperatorName:    op.FirstName + " " + op.LastName,
		Permissions:     op.Permissions,
		Preferences:     pref,
	})

	return ses, nil
}

func split(mail string) (string, string) {
	arr := strings.Split(mail, "@")
	if len(arr) < 2 {
		return "", ""
	}
	return arr[0], arr[1]
}
