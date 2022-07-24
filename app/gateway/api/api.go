package api

import (
	"comies/app"
	"comies/app/gateway/api/handler"
	"comies/app/gateway/api/menu"
	"comies/app/gateway/api/middleware"
	"comies/app/gateway/api/ordering"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

var Logger zap.Logger

func NewAPI(application app.Application) chi.Router {

	mdl := middleware.NewMiddlewares(application.Managers)
	h := handler.NewHandler(map[string]handler.Middleware{
		"tx": mdl.Transaction,
	})

	r := chi.NewRouter().With(mdl.Logging, mdl.CORS)

	r.Route("/menu", func(r chi.Router) {
		r = h.RegisterService(r, menu.NewService(application.Menu))
	})

	r.Route("/ordering", func(r chi.Router) {
		r = h.RegisterService(r, ordering.NewService(application.Ordering))
	})

	return r
}
