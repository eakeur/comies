package products

import "comies/app/jobs/menu"

type Handler struct {
	products menu.Jobs
}

func NewHandler(j menu.Jobs) Handler {
	return Handler{
		products: j,
	}
}
