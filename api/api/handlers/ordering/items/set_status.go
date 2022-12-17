package items

import (
	"comies/api/request"
	"comies/api/send"
	"comies/core/types"
	"context"
	"net/http"
)

func (h Handler) SetStatus(ctx context.Context, r request.Request) send.Response {
	id, err := r.IDParam("item_id")
	if err != nil {
		return send.IDError(err)
	}

	var st SetItemStatusRequest
	err = r.JSONBody(&st)
	if err != nil {
		return send.JSONError(err)
	}

	err = h.items.SetItemStatus(ctx, id, st.Status)
	if err != nil {
		return send.FromError(err)
	}

	r.Commit(ctx)
	return send.Data(http.StatusNoContent, nil)
}

type SetItemStatusRequest struct {
	Status types.Status `json:"status"`
}
