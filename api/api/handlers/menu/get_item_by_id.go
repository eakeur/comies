package menu

import (
	"comies/api/request"
	"comies/api/send"
	"context"
	"net/http"
)

// GetProductByKey fetches a product by its ID or code.
//
// @Summary     Fetches a product
// @Description Fetches a product by one of itd unique keys (id or code).
// @Tags        Product
// @Param       product_key path     string false "The product ID"
// @Param       code        query    bool   false "Toggles if the API should search by code"
// @Success     200         {object} rest.Response{data=GetProductByKeyResponse{}}
// @Failure     404         {object} rest.Response{error=rest.Error{}} "PRODUCT_NOT_FOUND"
// @Failure     400         {object} rest.Response{error=rest.Error{}} "INVALID_ID"
// @Failure     500         {object} rest.Response{error=rest.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /menu/products/{product_id} [GET]
func (h Handler) GetItemByID(ctx context.Context, r request.Request) send.Response {
	id, err := r.IDParam(ItemIDParam)
	if err != nil {
		return send.IDError(err)
	}

	p, err := h.menu.GetProductByID(ctx, id)
	if err != nil {
		return send.FromError(err)
	}

	return send.Data(http.StatusOK, Product{
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
}
