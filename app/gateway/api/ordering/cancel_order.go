package ordering

import (
	"comies/app/gateway/api/failures"
	"comies/app/gateway/api/handler"
	"comies/app/sdk/throw"
	"context"
	"net/http"
)

func (s Service) CancelOrder(ctx context.Context, r *http.Request) handler.Response {
	id, err := handler.GetResourceIDFromURL(r, "order_id")
	if err != nil {
		return handler.IDParsingErrorResponse(err)
	}

	err = s.ordering.CancelOrder(ctx, id)
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	return handler.ResponseWithData(http.StatusNoContent, nil)
}
