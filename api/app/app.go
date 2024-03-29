package app

import (
	"comies/core/types"
	"comies/data/repos/billing/bill"
	bill_item "comies/data/repos/billing/item"
	"comies/data/repos/menu/ingredient"
	"comies/data/repos/menu/movement"
	"comies/data/repos/menu/price"
	"comies/data/repos/menu/product"
	"comies/data/repos/ordering/item"
	"comies/data/repos/ordering/order"
	"comies/data/repos/ordering/status"
	"comies/jobs/billing"
	"comies/jobs/menu"
	"comies/jobs/ordering"

	"github.com/bwmarrin/snowflake"
)

type Deps struct {
	Snowflake *snowflake.Node
}

func NewApp(deps Deps) App {
	repos := repositories()

	idCreator := func() types.ID {
		return types.ID(deps.Snowflake.Generate())
	}

	menu := menu.NewJobs(menu.Deps{
		Products:    repos.Menu.Products,
		Ingredients: repos.Menu.Ingredients,
		Movements:   repos.Menu.Movements,
		Prices:      repos.Menu.Prices,
		IDCreator:   idCreator,
	})

	billing := billing.NewJobs(billing.Deps{
		Bills:     repos.Billing.Bill,
		Items:     repos.Billing.Item,
		IDCreator: idCreator,
	})

	ordering := ordering.NewJobs(ordering.Deps{
		Orders:    repos.Ordering.Orders,
		Items:     repos.Ordering.Items,
		Statuses:  repos.Ordering.Statuses,
		IDCreator: idCreator,
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
