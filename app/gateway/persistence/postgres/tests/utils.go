package tests

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/app/core/entities/ingredient"
	"gomies/app/core/entities/item"
	"gomies/app/core/entities/movement"
	"gomies/app/core/entities/order"
	"gomies/app/core/entities/product"
	"gomies/app/core/entities/stock"
)

func (d *Database) CheckValue(ctx context.Context, script string, expected interface{}, args ...interface{}) {
	var got interface{}
	r := d.Pool.QueryRow(ctx, script, args...)
	if err := r.Scan(&got); err != nil {
		d.Test.Errorf("an error occurred when checking value: %v", err)
	}
	assert.Equal(d.Test, expected, got)
}

func (d *Database) InsertOrders(ctx context.Context, orders ...order.Order) ([]order.Order, error) {
	const script = `
		insert into orders (
			id,
			identification, 
			placed_at,
			delivery_mode,
			observations,
			address,
			phone
		) values (
			$1, $2, $3, $4, $5, $6, $7
		)
	`

	for _, o := range orders {
		_, err := d.Pool.Exec(ctx, script,
			o.ID,
			o.Identification,
			o.PlacedAt,
			o.DeliveryMode,
			o.Observations,
			o.Address,
			o.Phone,
		)
		if err != nil {
			return nil, err
		}
	}

	return orders, nil
}

func (d *Database) InsertOrdersFlow(ctx context.Context, orders ...order.FlowUpdate) ([]order.FlowUpdate, error) {
	const script = `
		insert into orders_flow (
			id,
			order_id, 
			occurred_at,
			status
		) values (
			$1, $2, $3, $4
		)
	`

	for _, o := range orders {
		_, err := d.Pool.Exec(ctx, script,
			o.ID,
			o.OrderID,
			o.OccurredAt,
			o.Status,
		)
		if err != nil {
			return nil, err
		}
	}

	return orders, nil
}

func (d *Database) InsertItems(ctx context.Context, items ...item.Item) ([]item.Item, error) {
	const script = `
		insert into items (
			id,
			order_id,
			status,
            price,
			product_id,
			quantity,
			observations
		) values (
			$1, $2, $3, $4, $5, $6, $7
		)
	`

	for _, i := range items {
		_, err := d.Pool.Exec(ctx, script,
			i.ID,
			i.OrderID,
			i.Status,
			i.Price,
			i.ProductID,
			i.Quantity,
			i.Observations,
		)
		if err != nil {
			return nil, err
		}
	}

	return items, nil
}

func (d *Database) InsertProducts(ctx context.Context, products ...product.Product) ([]product.Product, error) {
	const script = `
		insert into products (
			id,
			code,
			name,
			type,
			cost_price,
			sale_price,
			sale_unit,
			minimum_sale
		) values (
			$1, $2, $3, $4, $5, $6, $7, $8
		)
	`

	for _, p := range products {
		_, err := d.Pool.Exec(ctx, script,
			p.ID,
			p.Active,
			p.Code,
			p.Name,
			p.Type,
			p.CostPrice,
			p.SalePrice,
			p.SaleUnit,
			p.MinimumSale,
		)
		if err != nil {
			return nil, err
		}
	}

	return products, nil
}

func (d *Database) InsertIngredients(ctx context.Context, ingredients ...ingredient.Ingredient) ([]ingredient.Ingredient, error) {
	const script = `
		insert into ingredients (
			id,
			product_id,
			ingredient_id,
			quantity,
			optional
		) values (
			$1, $2, $3, $4, $5
		)
	`

	for _, o := range ingredients {
		_, err := d.Pool.Exec(ctx, script,
			o.ID,
			o.ProductID,
			o.IngredientID,
			o.Quantity,
			o.Optional,
		)
		if err != nil {
			return nil, err
		}
	}

	return ingredients, nil
}

func (d *Database) InsertMovements(ctx context.Context, movements ...movement.Movement) ([]movement.Movement, error) {
	const script = `
		insert into movements (
			id,
			stock_id,
			type,
			date,
			quantity,
			value,
			agent_id
		) values (
			$1, $2, $3, $4, $5, $6, $7
		)
	`

	for _, o := range movements {
		_, err := d.Pool.Exec(ctx, script,
			o.ID,
			o.StockID,
			o.Type,
			o.Date,
			o.Quantity,
			o.PaidValue,
			o.AgentID,
		)
		if err != nil {
			return nil, err
		}
	}

	return movements, nil
}

func (d *Database) InsertStocks(ctx context.Context, stocks ...stock.Stock) ([]stock.Stock, error) {
	const script = `
		insert into stocks (
			id,
			target_id, 
			maximum_quantity,
			minimum_quantity,
			location
		) values (
			$1, $2, $3, $4, $5
		)
	`

	for _, o := range stocks {
		_, err := d.Pool.Exec(ctx, script,
			o.ID,
			o.TargetID,
			o.MaximumQuantity,
			o.MinimumQuantity,
			o.Location,
		)
		if err != nil {
			return nil, err
		}
	}

	return stocks, nil
}
