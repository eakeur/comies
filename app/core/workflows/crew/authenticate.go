package crew

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"gomies/app/core/entities/crew"
	"gomies/app/core/managers/session"
	"gomies/app/core/types/fault"
	"strings"
)

func (w workflow) Authenticate(ctx context.Context, auth crew.AuthRequest) (session.Session, error) {
	const operation = "Workflows.Operator.Authenticate"

	operatorNick, storeNick := split(auth.Nickname)

	op, err := w.crew.GetWithOperatorAndStoreNick(ctx, operatorNick, storeNick)
	if err != nil {
		return session.Session{}, fault.Wrap(err, operation)
	}

	err = bcrypt.CompareHashAndPassword([]byte(op.Password), []byte(auth.Password))
	if err != nil {
		return session.Session{}, crew.ErrWrongPassword
	}

	ctx, ses, err := w.sessions.Create(ctx, session.Session{
		OperatorID:      op.ExternalID,
		StoreExternalID: op.StoreExternalID,
		StoreInternalID: op.StoreID,
		OperatorName:    op.FirstName + " " + op.LastName,
		// TODO add preferences and permissions here
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
