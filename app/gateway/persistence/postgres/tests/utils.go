package tests

import (
	"context"
	"gomies/app/core/entities/ingredient"
	"gomies/app/core/entities/item"
	"gomies/app/core/entities/order"
	"gomies/app/core/entities/product"
)

func (d *Database) InsertOrders(ctx context.Context, orders ...order.Order) ([]order.Order, error) {
	const script = `
		insert into orders (
			id,
			store_id, 
			identification, 
			placed_at, 
			status,
			delivery_mode,
			observations
		) values (
			$1, $2, $3, $4, $5, $6, $7
		)
	`

	for _, o := range orders {
		_, err := d.Pool.Exec(ctx, script,
			o.ID,
			o.StoreID,
			o.Identification,
			o.PlacedAt,
			o.Status,
			o.DeliveryMode,
			o.Observations,
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
			observations,
			store_id
		) values (
			$1, $2, $3, $4, $5, $6, $7, $8
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
			i.StoreID,
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
			active,
			code,
			name,
			type,
			cost_price,
			sale_price,
			sale_unit,
			minimum_sale,
			store_id
		) values (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10
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
			p.StoreID,
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
			optional,
			store_id
		) values (
			$1, $2, $3, $4, $5, $6
		)
	`

	for _, o := range ingredients {
		_, err := d.Pool.Exec(ctx, script,
			o.ID,
			o.ProductID,
			o.IngredientID,
			o.Quantity,
			o.Optional,
			o.StoreID,
		)
		if err != nil {
			return nil, err
		}
	}

	return ingredients, nil
}
