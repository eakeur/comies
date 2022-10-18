package route

import (
	"comies/app/api/request"
	"comies/app/api/send"
	"comies/app/core/types"
	"context"
	"net/http"
)

type Route func(ctx context.Context, r request.Request) send.Response

func IDTargetRoute(idParamName string, fn func(ctx context.Context, i types.ID) error) Route {
	return func(ctx context.Context, r request.Request) send.Response {
		i, err := r.IDParam(idParamName)
		if err != nil {
			return send.IDError(err)
		}

		err = fn(ctx, i)
		if err != nil {
			return send.FromError(err)
		}

		return send.Data(http.StatusNoContent, nil)
	}
}
