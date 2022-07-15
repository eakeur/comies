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
	"github.com/jackc/pgx/v4/pgxpool"
)

type (
	Gateways struct {
		database *pgxpool.Pool
	}

	Actions struct {
		products    product.Actions
		ingredients ingredient.Actions
		stocks      stock.Actions
		movements   movement.Actions
		orders      order.Actions
		items       item.Actions
	}

	Services struct {
		products ordering.MenuService
		stocks   menu.StockService
	}

	Application struct {
		menu     menu.Workflow
		ordering ordering.Workflow
		stocking stocking.Workflow
	}
)
