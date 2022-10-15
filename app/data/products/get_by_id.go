package products

import (
	"comies/app/core/id"
	"comies/app/core/menu"
	"comies/app/data/conn"
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
)

func GetByID(ctx context.Context, id id.ID) (menu.Product, error) {
	const script = `
		select
			p.id,
			p.code,
			p.name,
			p.type,
			p.cost_price,
			p.sale_price,
			p.sale_unit,
			p.minimum_sale,
			p.minimum_quantity,
			p.maximum_quantity,
			p.location,
			coalesce(m.balance, 0) as balance
		from
			products p
			left join products_balances m on p.id = m.product_id
		where
			p.id = $1
	`

	row, err := conn.QueryRowFromContext(ctx, script, id)
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
