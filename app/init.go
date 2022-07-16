package app

import (
	"comies/app/core/workflows/menu"
	"comies/app/core/workflows/ordering"
	"comies/app/core/workflows/stocking"
	"comies/app/gateway/persistence/postgres/ingredient"
	"comies/app/gateway/persistence/postgres/item"
	"comies/app/gateway/persistence/postgres/movement"
	"comies/app/gateway/persistence/postgres/order"
	"comies/app/gateway/persistence/postgres/product"
	"comies/app/gateway/persistence/postgres/stock"
	"comies/app/gateway/persistence/postgres/transaction"
	stockingService "comies/app/gateway/services/menu/stocking"
	menuService "comies/app/gateway/services/ordering/menu"
	"comies/app/sdk/id"
)

func NewApplication(gateways Gateways) Application {

	var (
		actions = Actions{
			Products:    product.NewActions(gateways.Database),
			Ingredients: ingredient.NewActions(gateways.Database),
			Stocks:      stock.NewActions(gateways.Database),
			Movements:   movement.NewActions(gateways.Database),
			Orders:      order.NewActions(gateways.Database),
			Items:       item.NewActions(gateways.Database),
		}

		managers = Managers{
			Transactions: transaction.NewManager(gateways.Database),
			ID:           id.NewManager(gateways.SnowflakeNode),
		}

		services = Services{}
	)

	stocks := stocking.NewWorkflow(actions.Stocks, actions.Movements, managers.ID)
	services.Stocks = stockingService.NewService(stocks)

	menus := menu.NewWorkflow(actions.Products, actions.Ingredients, services.Stocks, managers.ID)
	services.Products = menuService.NewService(menus)

	orders := ordering.NewWorkflow(actions.Orders, actions.Items, services.Products, managers.ID)

	return Application{
		Managers: managers,
		Menu:     menus,
		Ordering: orders,
		Stocking: stocks,
	}
}
