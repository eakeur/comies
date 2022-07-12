package member

import (
	"context"
	"gomies/app/core/entities/member"
	"gomies/app/gateway/persistence/postgres/query"
	"gomies/app/sdk/fault"
)

func (a actions) List(ctx context.Context, filter member.Filter) ([]member.Member, int, error) {
	const script = `
		select
			m.id,
			m.active,
			m.nickname,
			m.password
		from
			member m
		%where_query%
	`

	q := query.NewQuery(script).
		Where(filter.Name != "", "m.name = $%v", filter.Name)

	rows, err := a.db.Query(ctx, q.Script(), q.Args)
	if err != nil {
		return nil, 0, fault.Wrap(err)
	}

	members := make([]member.Member, 0)
	for rows.Next() {
		var m member.Member
		err := rows.Scan(
			&m.ID,
			&m.Active,
			&m.Nickname,
			&m.Password,
		)
		if err != nil {
			continue
		}

		members = append(members, m)
	}

	return members, 0, nil
}
