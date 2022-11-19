package api

import (
	"comies/app/api/handlers/v1/menu"
	"comies/app/api/route"
	"comies/app/config"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Middlewares struct {
	Logging, CORS, TX, Pool func(http.Handler) http.Handler
}

func Serve(cfg config.Config, middlewares Middlewares) error {
	r := chi.NewRouter().With(middlewares.Logging, middlewares.CORS)

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/products", func(r chi.Router) {
			tx := r.With(middlewares.Pool)
			tx.Post("/", route.Route(menu.CreateProduct))
			tx.Put("/{product_id}", route.Route(menu.UpdateProduct))

			pool := r.With(middlewares.TX)
			pool.Post("/", route.Route(menu.ListProducts))
			pool.Get("/{product_id}", route.Route(menu.GetProductByKey))
		})

		r.Route("/ordering", func(r chi.Router) {

		})

		r.Route("/realtime", func(r chi.Router) {

		})
	})

	return nil
}
