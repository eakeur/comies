package movements

import (
	"comies/app/api/request"
	"comies/app/api/send"
	"comies/app/core/menu/movement"
	"context"
	"net/http"
	"time"
)

// GetProductMovements fetches all product movements.
//
// @Summary     Fetches movements
// @Description Fetches all product movements.
// @Tags        Product
// @Param       product_id path     string false "The product ID"
// @Param       start       query    string false "Adds a filter looking for the start date"
// @Param       end         query    string false "Adds a filter looking for the end date"
// @Success     200         {object} rest.Response{data=[]ListMovementsResponse{}}
// @Failure     500         {object} rest.Response{error=rest.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /menu/products/{product_id}/movements [GET]
func (h Handler) List(ctx context.Context, r request.Request) send.Response {
	id, err := r.IDParam("product_id")
	if err != nil {
		return send.IDError(err)
	}

	var filter movement.Filter
	filter.ProductID = id

	var query = r.URL.Query()
	if parse, err := time.Parse(time.RFC3339, query.Get("start")); err == nil {
		filter.InitialDate = parse
	}

	if parse, err := time.Parse(time.RFC3339, query.Get("end")); err == nil {
		filter.FinalDate = parse
	}

	list, err := h.movements.ListMovements(ctx, filter)
	if err != nil {
		return send.FromError(err)
	}

	return send.Data(http.StatusOK, list)
}
