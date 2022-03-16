package wrappers

import (
	"gomies/app/core/entities/category"
	"gomies/app/core/entities/contacting"
	"gomies/app/core/entities/crew"
	"gomies/app/core/entities/product"
	"gomies/app/core/entities/stock"
	"gomies/app/core/entities/store"
)

type Actions struct {
	Products   product.Actions
	Stocks     stock.Actions
	Categories category.Actions
	Contacting contacting.Actions
	Stores     store.Actions
	Crew       crew.Actions
}
