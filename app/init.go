package app

import (
	"comies/app/core/workflows/menu"
	"comies/app/core/workflows/ordering"
	"comies/app/gateway/persistence/postgres/ingredient"
	"comies/app/gateway/persistence/postgres/item"
	"comies/app/gateway/persistence/postgres/movement"
	"comies/app/gateway/persistence/postgres/order"
	"comies/app/gateway/persistence/postgres/product"
	"comies/app/gateway/persistence/postgres/transaction"
	"comies/app/sdk/id"
)

func NewApplication(gateways Gateways) Application {

	var (
		actions = Actions{
			Products:    product.NewActions(gateways.Database),
			Ingredients: ingredient.NewActions(gateways.Database),
			Movements:   movement.NewActions(gateways.Database),
			Orders:      order.NewActions(gateways.Database),
			Items:       item.NewActions(gateways.Database),
		}

		managers = Managers{
			Logger:       gateways.Logger,
			Transactions: transaction.NewManager(gateways.Database),
			ID:           id.NewManager(gateways.SnowflakeNode),
		}
	)

	menus := menu.NewWorkflow(actions.Products, actions.Ingredients, actions.Movements, managers.ID)

	orders := ordering.NewWorkflow(actions.Orders, actions.Items, menus, managers.ID)

	return Application{
		Managers: managers,
		Menu:     menus,
		Ordering: orders,
	}
}
