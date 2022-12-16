package menu

import (
	"comies/api/request"
	"comies/api/send"
	"comies/core/menu/product"
	"context"
	"net/http"
)

// UpdateItem updates a product to the store's menu.
//
// @Summary     Updates product
// @Description updates a product to the store's menu.
// @Tags        Product
// @Param       product body UpdateProductRequest true "The properties to define the product"
// @Success     204
// @Failure     400 {object} rest.Response{error=rest.Error{}} "INVALID_ID"
// @Failure     412     {object} rest.Response{error=rest.Error{}} "PRODUCT_CODE_ALREADY_EXISTS"
// @Failure     422     {object} rest.Response{error=rest.Error{}} "PRODUCT_ZERO_SALE_QUANTITY, PRODUCT_ZERO_PRICE, PRODUCT_INVALID_CODE, PRODUCT_INVALID_NAME, PRODUCT_INVALID_TYPE"
// @Failure     500     {object} rest.Response{error=rest.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /menu/products/{product_id} [PUT]
func (h Handler) UpdateItem(ctx context.Context, r request.Request) send.Response {

	id, err := r.IDParam(ItemIDParam)
	if err != nil {
		return send.IDError(err)
	}

	var p Item
	err = r.JSONBody(&p)
	if err != nil {
		return send.JSONError(err)
	}

	err = h.menu.UpdateProduct(ctx, product.Product{
		ID:              id,
		Code:            p.Code,
		Name:            p.Name,
		Type:            p.Type,
		CostPrice:       p.CostPrice,
		SaleUnit:        p.SaleUnit,
		MinimumSale:     p.MinimumSale,
		MaximumQuantity: p.MaximumQuantity,
		MinimumQuantity: p.MinimumQuantity,
		Location:        p.Location,
	})
	if err != nil {
		return send.FromError(err)
	}

	r.Commit(ctx)

	return send.Data(http.StatusNoContent, nil)

}
