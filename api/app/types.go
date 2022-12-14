package app

import (
	"comies/core/billing/bill"
	bill_item "comies/core/billing/item"
	"comies/core/menu/ingredient"
	"comies/core/menu/movement"
	"comies/core/menu/price"
	"comies/core/menu/product"
	"comies/core/ordering/item"
	"comies/core/ordering/order"
	"comies/core/ordering/status"
	"comies/jobs/billing"
	"comies/jobs/menu"
	"comies/jobs/ordering"
)

type App struct {
	Menu     menu.Jobs
	Ordering ordering.Jobs
	Billing  billing.Jobs
}

type Repositories struct {
	Menu     MenuRepositories
	Ordering OrderingRepositories
	Billing  BillingRepositories
}

type MenuRepositories struct {
	Products    product.Actions
	Ingredients ingredient.Actions
	Movements   movement.Actions
	Prices      price.Actions
}

type OrderingRepositories struct {
	Orders   order.Actions
	Items    item.Actions
	Statuses status.Actions
}

type BillingRepositories struct {
	Bill bill.Actions
	Item bill_item.Actions
}
