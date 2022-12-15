package orders

import "comies/jobs/ordering"

type Handler struct {
	orders ordering.Jobs
}

func NewHandler(j ordering.Jobs) Handler {
	return Handler{
		orders: j,
	}
}
