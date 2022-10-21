package menu

import (
	"comies/app/api/request"
	"comies/app/api/send"
	"comies/app/core/product"
	"comies/app/core/types"
	"comies/app/data/products"
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
func GetProductByKey(ctx context.Context, r request.Request) send.Response {
	key := r.Param("product_key")

	var p product.Product
	var err error

	switch r.GetQuery("code") {
	case "true":
		p, err = products.GetByCode(ctx, key)
	default:
		i, err := types.FromString(key)
		if err != nil {
			return send.IDError(err)
		}

		p, err = products.GetByID(ctx, i)
	}

	if err != nil {
		return send.FromError(err)
	}

	return send.Data(http.StatusOK, p)
}
