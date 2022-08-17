package product

import (
	"comies/app/core/entities/product"
	"comies/app/core/throw"
	"comies/app/gateway/persistence/postgres/query"
	"context"
	"fmt"
	"strings"
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
			p.location,
			coalesce(m.balance, 0) as balance
		from
			products p
			left join products_balances m on p.id = m.product_id
		%where_query%
		%order%
	`

	q := query.NewQuery(script).
		Where(filter.Code != "", "p.code like $%v", filter.Code+"%").And().
		Where(filter.Name != "", "p.name like $%v", "%"+filter.Name+"%").And().
		Where(filter.Type != 0, "p.type = $%v", filter.Type)

	scr := fmt.Sprintf(strings.Replace(q.Script(), "%order%", `
	order by 
			p.type in ($%d, $%d),
			coalesce(m.balance, 0) - p.minimum_quantity, 
			p.code
	`, 1), len(q.Args)+1, len(q.Args)+2)

	rows, err := a.db.Query(ctx, scr, append(q.Args, product.InputType, product.OutputType)...)
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
			&p.CostPrice,
			&p.SalePrice,
			&p.SaleUnit,
			&p.MinimumSale,
			&p.MinimumQuantity,
			&p.MaximumQuantity,
			&p.Location,
			&p.Balance,
		); err != nil {
			continue
		}

		products = append(products, p)
	}

	return products, nil
}
