package items

import "comies/jobs/ordering"

type Handler struct {
	items ordering.Jobs
}

func NewHandler(j ordering.Jobs) Handler {
	return Handler{
		items: j,
	}
}
