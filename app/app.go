package app

import (
	"comies/app/core/entities/ingredient"
	"comies/app/core/entities/item"
	"comies/app/core/entities/movement"
	"comies/app/core/entities/order"
	"comies/app/core/entities/product"
	"comies/app/core/entities/stock"
	"comies/app/core/workflows/menu"
	"comies/app/core/workflows/ordering"
	"comies/app/core/workflows/stocking"
	"comies/app/gateway/persistence/postgres/transaction"
	"comies/app/sdk/id"

	"github.com/bwmarrin/snowflake"
	"github.com/jackc/pgx/v4/pgxpool"
)

type (
	Gateways struct {
		Database      *pgxpool.Pool
		SnowflakeNode *snowflake.Node
	}

	Actions struct {
		Products    product.Actions
		Ingredients ingredient.Actions
		Stocks      stock.Actions
		Movements   movement.Actions
		Orders      order.Actions
		Items       item.Actions
	}

	Services struct {
		Products ordering.MenuService
		Stocks   menu.StockService
	}

	Managers struct {
		Transactions transaction.Manager
		ID           id.Manager
	}

	Application struct {
		Managers Managers
		Menu     menu.Workflow
		Ordering ordering.Workflow
		Stocking stocking.Workflow
	}
)
