package ingredients

import "comies/jobs/menu"

type Handler struct {
	ingredients menu.Jobs
}

func NewHandler(j menu.Jobs) Handler {
	return Handler{
		ingredients: j,
	}
}
