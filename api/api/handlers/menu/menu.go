package menu

import "comies/jobs/menu"

const (
	ItemIDParam       = "item_id"
	IngredientIDParam = "ingredient_id"
	MovementIDParam   = "movement_id"
	PriceParam        = "price"
)

type Handler struct {
	menu menu.Jobs
}

func NewHandler(j menu.Jobs) Handler {
	return Handler{
		menu: j,
	}
}
