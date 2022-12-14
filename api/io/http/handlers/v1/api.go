package v1

import (
	"comies/app"
	"comies/io/http/handlers/v1/menu/ingredients"
	"comies/io/http/handlers/v1/menu/movements"
	"comies/io/http/handlers/v1/menu/prices"
	"comies/io/http/handlers/v1/menu/products"
	"comies/io/http/handlers/v1/ordering/orders"
	"comies/io/http/handlers/v1/ordering/statuses"
	"comies/io/http/route"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Dependencies struct {
	App      app.App
	TX, Pool func(http.Handler) http.Handler
}

func Serve(router chi.Router, deps Dependencies) {
	router.
		Route("/api/v1", func(r chi.Router) {

			r.Route("/menu", func(r chi.Router) {
				r.Route("/products", func(r chi.Router) {
					p := products.NewHandler(deps.App.Menu)

					r.With(deps.TX).Post("/", route.Route(p.Create))
					r.With(deps.Pool).Get("/", route.Route(p.List))

					r.Route("/{product_id}", func(r chi.Router) {
						r.With(deps.TX).Put("/", route.Route(p.Update))
						r.With(deps.Pool).Get("/", route.Route(p.GetByID))
						r.With(deps.Pool).Get("/name", route.Route(p.GetNameByID))

						r.Route("/ingredients", func(r chi.Router) {
							i := ingredients.NewHandler(deps.App.Menu)

							r.With(deps.TX).Post("/", route.Route(i.Create))
							r.With(deps.TX).Delete("/{ingredient_id}", route.Route(i.Remove))
							r.With(deps.Pool).Get("/", route.Route(i.List))
						})

						r.Route("/prices", func(r chi.Router) {
							p := prices.NewHandler(deps.App.Menu)

							r.With(deps.Pool).Get("/", route.Route(p.List))
							r.With(deps.Pool).Get("/latest", route.Route(p.Latest))
							r.With(deps.TX).Post("/{value}", route.Route(p.Create))
						})

						r.Route("/movements", func(r chi.Router) {
							m := movements.NewHandler(deps.App.Menu)

							r.With(deps.TX).Post("/", route.Route(m.Create))
							r.With(deps.TX).Delete("/{movement_id}", route.Route(m.Remove))
							r.With(deps.Pool).Get("/balance", route.Route(m.GetProductStockBalance))
							r.With(deps.Pool).Get("/", route.Route(m.List))
						})
					})
				})
			})

			r.Route("/ordering", func(r chi.Router) {
				r.Route("/orders", func(r chi.Router) {
					o := orders.NewHandler(deps.App.Ordering)

					r.With(deps.TX).Post("/", route.Route(o.Place))
					r.With(deps.Pool).Get("/", route.Route(o.List))

					r.Route("/{order_id}", func(r chi.Router) {
						r.With(deps.Pool).Get("/", route.Route(o.GetCustomer))
						r.With(deps.TX).Delete("/", route.Route(o.Cancel))

						r.Route("/status", func(r chi.Router) {
							s := statuses.NewHandler(deps.App.Ordering)

							r.With(deps.Pool).Get("/", route.Route(s.Get))
							r.With(deps.TX).Put("/{status}", route.Route(s.Set))
						})
					})
				})
			})
		})
}