package product

import (
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
)

func (a actions) GetNameByID(ctx context.Context, id types.ID) (string, error) {
	const script = `
		select p.name from products p where p.id = $1
	`

	row := a.db.QueryRow(ctx, script, id)

	var name string
	if err := row.Scan(&name); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", throw.Error(throw.ErrNotFound).
				Describe("the product id provided seems to not exist").Params(map[string]interface{}{
				"id": id,
			})
		}
		return "", throw.Error(err)
	}

	return name, nil
}
