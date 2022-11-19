package movements

import (
	"comies/app/api/request"
	"comies/app/api/send"
	"context"
	"net/http"
)

func (h Handler) Remove(ctx context.Context, r request.Request) send.Response {

	id, err := r.IDParam("id")
	if err != nil {
		return send.IDError(err)
	}

	err = h.movements.RemoveMovement(ctx, id)
	if err != nil {
		return send.FromError(err)
	}

	r.Commit(ctx)

	return send.Data(http.StatusNoContent, nil)
}
