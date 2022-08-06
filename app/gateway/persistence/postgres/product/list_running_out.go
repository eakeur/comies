package product

import (
	"comies/app/core/entities/product"
	"comies/app/core/throw"
	"context"
)

func (a actions) ListRunningOut(ctx context.Context) ([]product.Product, error) {
	const script = `
		select
		p.id,
		p.code,
		p.name,
		p.type,
		p.sale_unit,
		p.minimum_quantity,
		p.maximum_quantity,
		coalesce(m.balance, 0) as balance
	from
		products p
		left join (
			select
				m.product_id,
				sum(m.quantity) as balance
			from movements m
			group by m.product_id
		) m on p.id = m.product_id
	where
		coalesce(m.balance, 0) <= (p.maximum_quantity * 0.25)
		and p.type in ($1, $2)
	order by 
		coalesce(m.balance, 0) - p.minimum_quantity
	`

	rows, err := a.db.Query(ctx, script, product.OutputType, product.InputType)
	if err != nil {
		return nil, throw.Error(err)
	}

	products := make([]product.Product, 0)
	for rows.Next() {
		var p product.Product
		if err := rows.Scan(
			&p.ID,
			&p.Code,
			&p.Name,
			&p.Type,
			&p.SaleUnit,
			&p.MinimumQuantity,
			&p.MaximumQuantity,
			&p.Balance,
		); err != nil {
			continue
		}

		products = append(products, p)
	}

	return products, nil
}
