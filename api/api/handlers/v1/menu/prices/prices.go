package prices

import "comies/jobs/menu"

type Handler struct {
	prices menu.Jobs
}

func NewHandler(j menu.Jobs) Handler {
	return Handler{
		prices: j,
	}
}
