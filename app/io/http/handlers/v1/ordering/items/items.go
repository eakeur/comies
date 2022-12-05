package items

import "comies/app/jobs/ordering"

type Handler struct {
	items ordering.Jobs
}

func NewHandler(j ordering.Jobs) Handler {
	return Handler{
		items: j,
	}
}
