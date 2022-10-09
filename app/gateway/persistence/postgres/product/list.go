package product

import (
	"comies/app/core/entities/product"
	"comies/app/gateway/persistence/postgres/query"
	"context"
)

func (a actions) List(ctx context.Context, filter product.Filter) ([]product.Product, error) {
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
			p.location
		from
			products p
		%where_query%
		order by
			p.code
	`

	q := query.NewQuery(script).
		Where(filter.Code != "", "p.code like $%v", filter.Code+"%").And().
		Where(filter.Name != "", "p.name like $%v", "%"+filter.Name+"%").And().
		Where(filter.Type != 0, "p.type = $%v", filter.Type)

	rows, err := a.db.Query(ctx, q.Script(), q.Args...)
	if err != nil {
		return nil, err
	}

	products := make([]product.Product, 0)
	for rows.Next() {
		var p product.Product
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
		); err != nil {
			continue
		}

		products = append(products, p)
	}

	return products, nil
}
