package products

import (
	"comies/app/core/product"
	"comies/app/core/types"
	"comies/app/data/conn"
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
)

func GetNameByID(ctx context.Context, id types.ID) (string, error) {
	const script = `
		select p.name from products p where p.id = $1
	`

	row, err := conn.QueryRowFromContext(ctx, script, id)
	if err != nil {
		return "", err
	}

	var name string
	if err := row.Scan(&name); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", product.ErrNotFound
		}
		return "", err
	}

	return name, nil
}
