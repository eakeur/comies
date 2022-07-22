package api

import (
	"comies/app"
	"comies/app/gateway/api/handler"
	"comies/app/gateway/api/menu"
	"comies/app/gateway/api/middleware"
	"github.com/go-chi/chi/v5"
)

type (
	StatusCodeLoggingKey struct{}
)

func NewAPI(application app.Application) chi.Router {

	mdl := middleware.NewMiddlewares(application.Managers)
	h := handler.NewHandler(map[string]handler.Middleware{
		"tx": mdl.Transaction,
	})

	r := chi.NewRouter().With(mdl.Logging)

	r.Route("/menu", func(r chi.Router) {
		r = h.RegisterService(r, menu.NewService(application.Menu))
	})

	return r
}
