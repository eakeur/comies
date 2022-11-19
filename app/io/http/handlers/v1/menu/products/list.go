package products

import (
	"comies/app/api/request"
	"comies/app/api/send"
	"comies/app/core/menu/product"
	"comies/app/core/types"
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
func (h Handler) List(ctx context.Context, r request.Request) send.Response {
	query := r.URL.Query()

	values := query["types"]

	ty := make([]types.Type, len(values))
	for i, v := range values {
		idx, _ := strconv.Atoi(v)
		ty[i] = types.Type(idx)
	}

	filter := product.Filter{
		Code:  query.Get("code"),
		Name:  query.Get("name"),
		Types: ty,
	}

	list, err := h.products.ListProducts(ctx, filter)
	if err != nil {
		return send.FromError(err)
	}

	return send.Data(http.StatusOK, list)
}
