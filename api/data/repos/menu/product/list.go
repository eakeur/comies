package product

import (
	"comies/core/menu/product"
	"comies/data/conn"
	"comies/data/query"
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

	types := make([]interface{}, len(filter.Types))
	for i, t := range filter.Types {
		types[i] = t
	}

	q := query.NewQuery(script).
		Where(filter.Code != "", "p.code like $%v", filter.Code+"%").And().
		Where(filter.Name != "", "p.name like $%v", "%"+filter.Name+"%").And().
		Where(len(filter.Types) > 0, "p.type in (%v)", types...)

	rows, err := conn.QueryFromContext(ctx, q.Script(), q.Args...)
	if err != nil {
		return nil, err
	}

	return conn.ScanRows(rows,
		func(scan conn.Scan, p product.Product) error {
			return scan(
				&p.ID,
				&p.Code,
				&p.Name,
				&p.Type,
				&p.CostPrice,
				&p.SaleUnit,
				&p.MinimumSale,
				&p.MinimumQuantity,
				&p.MaximumQuantity,
				&p.Location,
			)
		},
	)
}
