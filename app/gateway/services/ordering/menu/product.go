package menu

import (
	"comies/app/core/workflows/menu"
	"comies/app/core/workflows/ordering"
)

var _ ordering.MenuService = service{}

type service struct {
	menu menu.Workflow
}

func NewService(menu menu.Workflow) ordering.MenuService {
	return service{menu: menu}
}
