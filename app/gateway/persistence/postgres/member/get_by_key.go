package member

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"gomies/app/core/entities/member"
	"gomies/app/gateway/persistence/postgres/query"
	"gomies/app/sdk/fault"
)

func (a actions) GetByKey(ctx context.Context, key member.Key) (member.Member, error) {
	const script = `
		select
			m.id,
        	m.active,
        	m.nickname,
        	m.password
		from
			members m
		%where_query%
	`

	q := query.NewQuery(script).
		Where(key.ID != 0, "m.id = $%v", key.ID).Or().
		Where(key.Nickname != "", "m.nickname = $%v", key.Nickname)

	row := a.db.QueryRow(ctx, q.Script(), q.Args)

	var m member.Member
	if err := row.Scan(
		&m.ID,
		&m.Active,
		&m.Nickname,
		&m.Password,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return member.Member{}, fault.Wrap(fault.ErrNotFound).
				Describe("the member id provided seems to not exist").Params(map[string]interface{}{
				"id":       key.ID,
				"nickname": key.Nickname,
			})
		}
		return member.Member{}, fault.Wrap(err)
	}
	return m, nil
}
