package menu

import (
	"comies/app/core/entities/movement"
	"comies/app/core/types"
	"comies/app/gateway/api/handler"
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
// @Success     200         {object} handler.Response{data=[]ListMovementsResponse{}}
// @Failure     500         {object} handler.Response{error=handler.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /menu/products/{product_id}/movements [GET]
func (s Service) GetProductMovements(ctx context.Context, r *http.Request) handler.Response {
	id, err := handler.GetResourceIDFromURL(r, "product_id")
	if err != nil {
		return handler.IDParsingErrorResponse(err)
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

	list, err := s.menu.ListMovements(ctx, filter)
	if err != nil {
		return handler.Fail(err)
	}

	return handler.ResponseWithData(http.StatusOK, NewListMovementsResponse(list))
}

type ListMovementsResponse struct {
	// ID is the unique identifier of this movement
	ID string `json:"id"`

	// ProductID is an identifier for the stock this movement references to
	ProductID string `json:"product_id"`

	// Type points out if this movement is input or output
	Type movement.Type `json:"type"`

	// Date is when the object got into the stock effectively
	Date time.Time `json:"date"`

	// Quantity is the amount being inserted or removed from this stock
	Quantity types.Quantity `json:"quantity"`

	// PaidValue is how much was paid/received for this resource
	PaidValue types.Currency `json:"paid_value"`

	// AgentID is the entity from this resource came from or is going to
	AgentID string `json:"agent_id"`
}

func NewListMovementsResponse(list []movement.Movement) []ListMovementsResponse {
	movements := make([]ListMovementsResponse, len(list))
	for i, p := range list {
		movements[i] = ListMovementsResponse{
			ID:        p.ID.String(),
			ProductID: p.ProductID.String(),
			Type:      p.Type,
			Date:      p.Date,
			Quantity:  p.Quantity,
			PaidValue: p.PaidValue,
			AgentID:   p.AgentID.String(),
		}
	}

	return movements
}
