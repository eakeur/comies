package ingredients

import "comies/app/jobs/menu"

type Handler struct {
	ingredients menu.Jobs
}

func NewHandler(j menu.Jobs) Handler {
	return Handler{
		ingredients: j,
	}
}
