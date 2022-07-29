package menu

import (
	"comies/app/core/workflows/menu"
	"comies/app/gateway/api/handler"
)

type Service struct {
	CreateProductRoute    handler.Route `path:"/products" method:"POST" body:"Product" middleware:"tx"`
	CreateMovementRoute   handler.Route `path:"/products/{product_id}/movements" method:"POST" body:"Movement" url:"product_id"`
	CreateIngredientRoute handler.Route `path:"/products/{product_id}/ingredients" method:"POST" body:"Ingredient" url:"product_id"`

	ListProductsRoute           handler.Route `path:"/products" method:"GET" params:"Filter"`
	GetProductByKeyRoute        handler.Route `path:"/products/{product_key}" method:"GET" url:"product_key"`
	GetProductNameByIDRoute     handler.Route `path:"/products/{product_id}/name" method:"GET" url:"product_id"`
	GetProductStockBalanceRoute handler.Route `path:"/products/{product_id}/stock-balance" method:"GET" url:"product_id"`
	GetProductMovementsRoute    handler.Route `path:"/products/{product_id}/movements" method:"GET" url:"product_id"`
	GetProductIngredientsRoute  handler.Route `path:"/products/{product_id}/ingredients" method:"GET" url:"product_id"`

	RemoveProductRoute           handler.Route `path:"/products/{product_id}" method:"DELETE" url:"product_id"`
	RemoveProductMovementRoute   handler.Route `path:"/products/{product_id}/movements/{movement_id}" method:"DELETE" url:"product_id"`
	RemoveProductIngredientRoute handler.Route `path:"/products/{product_id}/ingredients/{ingredient_id}" method:"DELETE" url:"product_id"`

	UpdateProductRoute handler.Route `path:"/products/{product_id}" method:"PUT" body:"Product" middleware:"tx"`

	menu menu.Workflow
}

func NewService(menu menu.Workflow) *Service {
	return &Service{
		menu: menu,
	}
}
