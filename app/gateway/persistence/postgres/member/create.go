package member

import (
	"context"
	"errors"
	"github.com/jackc/pgconn"
	"gomies/app/core/entities/member"
	"gomies/app/gateway/persistence/postgres"
	"gomies/app/gateway/persistence/postgres/transaction"
	"gomies/app/sdk/fault"
)

func (a actions) Create(ctx context.Context, m member.Member) (member.Member, error) {
	const script = `
		insert into members (
			id,
			active,
			name, 
			nickname,
			password
		) values (
			$1, $2, $3, $4, $5
		)
	`

	if _, err := transaction.ExecFromContext(ctx, script,
		m.ID,
		m.Active,
		m.Name,
		m.Nickname,
		m.Password,
	); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == postgres.DuplicateError && pgErr.ConstraintName == postgres.MemberIDPK {
				return member.Member{}, fault.Wrap(fault.ErrAlreadyExists).
					Describe("the member id provided seems to already exist").Params(map[string]interface{}{
					"id": m.ID,
				})
			}
			if pgErr.Code == postgres.DuplicateError && pgErr.ConstraintName == postgres.MemberNicknameUK {
				return member.Member{}, fault.Wrap(fault.ErrAlreadyExists).
					Describe("the member nickname provided seems to already exist").Params(map[string]interface{}{
					"id": m.ID,
				})
			}
		}

		return member.Member{}, err
	}

	return m, nil
}
