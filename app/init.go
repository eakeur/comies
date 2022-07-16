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
	stockingService "comies/app/gateway/services/menu/stocking"
	menuService "comies/app/gateway/services/ordering/menu"
)

func NewApplication(gateways Gateways) Application {

	var (
		actions = Actions{
			products:    product.NewActions(gateways.database),
			ingredients: ingredient.NewActions(gateways.database),
			stocks:      stock.NewActions(gateways.database),
			movements:   movement.NewActions(gateways.database),
			orders:      order.NewActions(gateways.database),
			items:       item.NewActions(gateways.database),
		}

		services = Services{}
	)

	stocks := stocking.NewWorkflow(actions.stocks, actions.movements)
	services.stocks = stockingService.NewService(stocks)

	menus := menu.NewWorkflow(actions.products, actions.ingredients, services.stocks)
	services.products = menuService.NewService(menus)

	orders := ordering.NewWorkflow(actions.orders, actions.items, services.products)

	return Application{
		Menu:     menus,
		Ordering: orders,
		Stocking: stocks,
	}
}
