package products

import (
	"comies/app/core/menu"
	"comies/app/data/conn"
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
)

func GetByCode(ctx context.Context, code string) (menu.Product, error) {
	const script = `
		select
			id,
			code,
			name,
			type,
			cost_price,
			sale_price,
			sale_unit,
			minimum_sale,
			minimum_quantity,
			maximum_quantity,
			location,
			coalesce(m.balance, 0) as balance
		from
			products p
		where
			p.code = $1
	`

	row, err := conn.QueryRowFromContext(ctx, script, code)
	if err != nil {
		return menu.Product{}, err
	}

	var p menu.Product
	if err := row.Scan(
		&p.ID,
		&p.Code,
		&p.Name,
		&p.Type,
		&p.CostPrice,
		&p.SalePrice,
		&p.SaleUnit,
		&p.MinimumSale,
		&p.MinimumQuantity,
		&p.MaximumQuantity,
		&p.Location,
		&p.Balance,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return menu.Product{}, menu.ErrNotFound
		}
		return menu.Product{}, err
	}

	return p, nil
}
