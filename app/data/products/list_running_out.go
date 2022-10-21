package products

import (
	"comies/app/core/product"
	"comies/app/data/conn"
	"context"
)

func ListRunningOut(ctx context.Context) ([]product.Product, error) {
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
		p.type in ($1, $2)
	order by 
		coalesce(m.balance, 0) - p.minimum_quantity
	`

	rows, err := conn.QueryFromContext(ctx, script, product.OutputType, product.InputType)
	if err != nil {
		return nil, err
	}

	products := make([]product.Product, 0)
	var p product.Product
	for rows.Next() {
		if err := rows.Scan(
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
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}
