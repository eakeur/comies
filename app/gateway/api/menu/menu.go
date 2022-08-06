package menu

import (
	"comies/app/core/workflows/menu"
	"comies/app/gateway/api/handler"
)

type Service struct {
	CreateProductRoute    handler.Route `path:"/products" method:"POST" middleware:"tx"`
	CreateMovementRoute   handler.Route `path:"/products/{product_id}/movements" method:"POST" middleware:"tx"`
	CreateIngredientRoute handler.Route `path:"/products/{product_id}/ingredients" method:"POST" middleware:"tx"`

	ListProductsRoute           handler.Route `path:"/products" method:"GET"`
	GetProductByKeyRoute        handler.Route `path:"/products/{product_key}" method:"GET"`
	GetProductNameByIDRoute     handler.Route `path:"/products/{product_id}/name" method:"GET"`
	GetProductStockBalanceRoute handler.Route `path:"/products/{product_id}/stock-balance" method:"GET"`
	GetProductMovementsRoute    handler.Route `path:"/products/{product_id}/movements" method:"GET"`
	GetProductIngredientsRoute  handler.Route `path:"/products/{product_id}/ingredients" method:"GET"`

	RemoveProductRoute           handler.Route `path:"/products/{product_id}" method:"DELETE" middleware:"tx"`
	RemoveProductMovementRoute   handler.Route `path:"/products/{product_id}/movements/{movement_id}" method:"DELETE" middleware:"tx"`
	RemoveProductIngredientRoute handler.Route `path:"/products/{product_id}/ingredients/{ingredient_id}" method:"DELETE" middleware:"tx"`

	UpdateProductRoute handler.Route `path:"/products/{product_id}" method:"PUT" middleware:"tx"`

	menu menu.Workflow
}

func NewService(menu menu.Workflow) *Service {
	return &Service{
		menu: menu,
	}
}
