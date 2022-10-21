package menu

import (
	"comies/app/api/request"
	"comies/app/api/send"
	"comies/app/core/product"
	"comies/app/core/types"
	"comies/app/data/products"
	"context"
	"net/http"
	"strconv"
)

// ListProducts fetches a product by its ID or code.
//
// @Summary     Fetches a product
// @Description Fetches a product by one of itd unique keys (id or code).
// @Tags        Product
// @Param       code query    string false "Adds a filter looking for the products codes"
// @Param       name query    string false "Adds a filter looking for the products names"
// @Param       type query    int    false "Adds a filter looking for the products types"
// @Success     200  {object} rest.Response{data=[]GetProductByKeyResponse{}}
// @Failure     500  {object} rest.Response{error=rest.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /menu/products [GET]
func ListProducts(ctx context.Context, r request.Request) send.Response {
	query := r.URL.Query()

	if query.Get("stock") == "true" {
		list, err := products.ListRunningOut(ctx)
		if err != nil {
			return send.FromError(err)
		}

		return send.Data(http.StatusOK, list)
	}

	ty, _ := strconv.Atoi(query.Get("type"))

	filter := product.Filter{
		Code: query.Get("code"),
		Name: query.Get("name"),
		Type: types.Type(ty),
	}

	list, err := products.List(ctx, filter)
	if err != nil {
		return send.FromError(err)
	}

	return send.Data(http.StatusOK, list)
}
