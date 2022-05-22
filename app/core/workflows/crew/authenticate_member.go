package crew

import (
	"context"
	"gomies/app/core/entities/member"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/session"
)

func (w workflow) AuthenticateMember(ctx context.Context, auth AuthRequest) (session.Session, error) {

	op, err := w.crew.GetByKey(ctx, member.Key{Nickname: auth.Nickname})
	if err != nil {
		return session.Session{}, fault.Wrap(err)
	}

	if err := op.Password.Compare(auth.Password); err != nil {
		return session.Session{}, err
	}

	_, ses, err := w.sessions.Create(ctx, session.Session{
		OperatorID:   op.ID,
		OperatorName: op.Name,
	})
	if err != nil {
		return session.Session{}, fault.Wrap(err)
	}

	return ses, nil
}
