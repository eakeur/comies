package handlers

import (
	"comies/api/handlers/menu"
	"comies/api/handlers/ordering/orders"
	"comies/api/handlers/ordering/statuses"
	"comies/api/route"
	"comies/app"
	"fmt"
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

			mn := menu.NewHandler(deps.App.Menu)
			r.Route("/menu", func(r chi.Router) {

				r.Route("/items", func(r chi.Router) {
					r.With(deps.TX).Post("/", route.Route(mn.CreateMenu))
					r.With(deps.Pool).Get("/", route.Route(mn.ListItems))

					r.Route(fmt.Sprintf("/{%s}", menu.ItemIDParam), func(r chi.Router) {
						r.With(deps.TX).Put("/", route.Route(mn.UpdateItem))
						r.With(deps.Pool).Get("/", route.Route(mn.GetItemByID))
						r.With(deps.Pool).Get("/name", route.Route(mn.GetItemNameByID))
						r.With(deps.Pool).Get("/stock", route.Route(mn.GetItemStockBalance))

						r.Route("/ingredients", func(r chi.Router) {
							r.With(deps.TX).Post("/", route.Route(mn.CreateItemIngredient))
							r.With(deps.TX).Delete(fmt.Sprintf("/{%s}", menu.IngredientIDParam), route.Route(mn.RemoveItemIngredient))
							r.With(deps.Pool).Get("/", route.Route(mn.ListItemIngredients))
						})

						r.Route("/prices", func(r chi.Router) {
							r.With(deps.TX).Post(fmt.Sprintf("/{%s}", menu.PriceParam), route.Route(mn.CreateItemPrice))
							r.With(deps.Pool).Get("/", route.Route(mn.ListItemPrices))
							r.With(deps.Pool).Get("/latest", route.Route(mn.GetLatestItemPrice))
						})

						r.Route("/movements", func(r chi.Router) {
							r.With(deps.TX).Post("/", route.Route(mn.CreateItemMovement))
							r.With(deps.TX).Delete(fmt.Sprintf("/{%s}", menu.MovementIDParam), route.Route(mn.RemoveItemMovement))
							r.With(deps.Pool).Get("/", route.Route(mn.ListItemMovements))
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
