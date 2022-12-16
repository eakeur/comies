package menu

import (
	"comies/api/request"
	"comies/api/send"
	"context"
	"net/http"
)

func (h Handler) RemoveItemMovement(ctx context.Context, r request.Request) send.Response {

	id, err := r.IDParam(MovementIDParam)
	if err != nil {
		return send.IDError(err)
	}

	err = h.menu.RemoveMovement(ctx, id)
	if err != nil {
		return send.FromError(err)
	}

	r.Commit(ctx)

	return send.Data(http.StatusNoContent, nil)
}
