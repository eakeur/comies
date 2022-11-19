package route

import (
	"comies/app/api/request"
	"comies/app/api/send"
	"comies/app/core/types"
	"comies/app/telemetry"
	"context"
	"net/http"
	"time"

	"go.uber.org/zap"
)

func Route(fn func(ctx context.Context, r request.Request) send.Response) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		now := time.Now().UTC()

		res := fn(ctx, request.Request{Request: r})

		log := telemetry.LoggerFromContext(ctx)

		if err := send.Write(w, res); err != nil {
			log.Error("failed writing response", zap.Error(err))
		}

		if err := res.Error; err != nil {
			log = log.With(zap.Error(err))
		}

		log.With(
			zap.Duration("took", time.Since(now)),
			zap.Int("status", res.Code),
		).Info("served request")
	})
}

func IDActionRoute(idParamName string, fn func(ctx context.Context, i types.ID) error) http.Handler {
	return Route(func(ctx context.Context, r request.Request) send.Response {
		i, err := r.IDParam(idParamName)
		if err != nil {
			return send.IDError(err)
		}

		err = fn(ctx, i)
		if err != nil {
			return send.FromError(err)
		}

		return send.Data(http.StatusNoContent, nil)
	})
}
