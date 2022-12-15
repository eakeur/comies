package products

import (
	"comies/api/request"
	"comies/api/send"
	"comies/core/menu/product"
	"comies/jobs/menu"
	"context"
)

// CreateProduct adds a product to the store's menu.
//
// @Summary     Create product
// @Description Adds a product to the store's menu.
// @Tags        Product
// @Param       product body     CreateProductRequest true "The properties to define the product"
// @Success     201     {object} rest.Response{data=CreateProductResponse{}}
// @Failure     412     {object} rest.Response{error=rest.Error{}} "PRODUCT_CODE_ALREADY_EXISTS"
// @Failure     422     {object} rest.Response{error=rest.Error{}} "PRODUCT_ZERO_SALE_QUANTITY, PRODUCT_ZERO_PRICE, PRODUCT_INVALID_CODE, PRODUCT_INVALID_NAME, PRODUCT_INVALID_TYPE"
// @Failure     500     {object} rest.Response{error=rest.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /menu/products [POST]
func (h Handler) Create(ctx context.Context, r request.Request) send.Response {
	var p Product
	err := r.JSONBody(&p)
	if err != nil {
		return send.JSONError(err)
	}

	id, err := h.products.CreateProduct(ctx, menu.ProductCreation{
		SalePrice: p.SalePrice,
		Product: product.Product{
			Code:            p.Code,
			Name:            p.Name,
			Type:            p.Type,
			CostPrice:       p.CostPrice,
			SaleUnit:        p.SaleUnit,
			MinimumSale:     p.MinimumSale,
			MaximumQuantity: p.MaximumQuantity,
			MinimumQuantity: p.MinimumQuantity,
			Location:        p.Location,
		},
	})
	if err != nil {
		return send.FromError(err)
	}

	r.Commit(ctx)

	return send.CreatedWithID(id)
}
