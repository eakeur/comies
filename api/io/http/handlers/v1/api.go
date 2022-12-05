package v1

import (
	"comies/io/http/handlers/v1/menu/ingredients"
	"comies/io/http/handlers/v1/menu/movements"
	"comies/io/http/handlers/v1/menu/prices"
	"comies/io/http/handlers/v1/menu/products"
	"comies/io/http/route"
	"comies/jobs/menu"
	"comies/jobs/ordering"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Dependencies struct {
	Menu     menu.Jobs
	Ordering ordering.Jobs
	TX, Pool func(http.Handler) http.Handler
}

func Serve(router chi.Router, deps Dependencies) {
	router.
		Route("/api/v1", func(r chi.Router) {

			r.Route("menu", func(r chi.Router) {
				r.Route("/products", func(r chi.Router) {
					p := products.NewHandler(deps.Menu)

					r.With(deps.TX).Post("/", route.Route(p.Create))
					r.With(deps.Pool).Get("/", route.Route(p.List))

					r.Route("/{product_id}", func(r chi.Router) {
						r.With(deps.TX).Put("/", route.Route(p.Update))
						r.With(deps.Pool).Get("/", route.Route(p.GetByID))
						r.With(deps.Pool).Get("/name", route.Route(p.GetNameByID))

						r.Route("/ingredients", func(r chi.Router) {
							i := ingredients.NewHandler(deps.Menu)

							r.With(deps.TX).Post("/", route.Route(i.Create))
							r.With(deps.TX).Delete("/{ingredient_id}", route.Route(i.Remove))
							r.With(deps.Pool).Get("/", route.Route(i.List))
						})

						r.Route("/prices", func(r chi.Router) {
							p := prices.NewHandler(deps.Menu)

							r.With(deps.Pool).Get("/", route.Route(p.List))
							r.With(deps.Pool).Get("/latest", route.Route(p.Latest))
							r.With(deps.TX).Post("/{value}", route.Route(p.Create))
						})

						r.Route("/movements", func(r chi.Router) {
							m := movements.NewHandler(deps.Menu)

							r.With(deps.TX).Post("/", route.Route(m.Create))
							r.With(deps.TX).Delete("/{movement_id}", route.Route(m.Remove))
							r.With(deps.Pool).Get("/balance", route.Route(m.GetProductStockBalance))
							r.With(deps.Pool).Get("/", route.Route(m.List))
						})
					})
				})
			})
		})
}
