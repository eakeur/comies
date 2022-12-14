package app

import (
	"comies/core/types"
	"comies/io/data/postgres/billing/bill"
	bill_item "comies/io/data/postgres/billing/item"
	"comies/io/data/postgres/menu/ingredient"
	"comies/io/data/postgres/menu/movement"
	"comies/io/data/postgres/menu/price"
	"comies/io/data/postgres/menu/product"
	"comies/io/data/postgres/ordering/item"
	"comies/io/data/postgres/ordering/order"
	"comies/io/data/postgres/ordering/status"
	"comies/jobs/billing"
	"comies/jobs/menu"
	"comies/jobs/ordering"
)

func NewApp(createID types.CreateID) App {
	repos := repositories()

	menu := menu.NewJobs(menu.Deps{
		Products:    repos.Menu.Products,
		Ingredients: repos.Menu.Ingredients,
		Movements:   repos.Menu.Movements,
		Prices:      repos.Menu.Prices,
		IDCreator:   createID,
	})

	billing := billing.NewJobs(billing.Deps{
		Bills: repos.Billing.Bill,
		Items: repos.Billing.Item,
	})

	ordering := ordering.NewJobs(ordering.Deps{
		Orders:    repos.Ordering.Orders,
		Items:     repos.Ordering.Items,
		Statuses:  repos.Ordering.Statuses,
		IDCreator: createID,
		Menu:      menu,
		Billing:   billing,
	})

	return App{
		Menu:     menu,
		Ordering: ordering,
	}
}

func repositories() Repositories {
	return Repositories{
		Menu: MenuRepositories{
			Products:    product.NewActions(),
			Movements:   movement.NewActions(),
			Ingredients: ingredient.NewActions(),
			Prices:      price.NewActions(),
		},
		Ordering: OrderingRepositories{
			Orders:   order.NewActions(),
			Items:    item.NewActions(),
			Statuses: status.NewActions(),
		},
		Billing: BillingRepositories{
			Bill: bill.NewActions(),
			Item: bill_item.NewActions(),
		},
	}
}
