package statuses

import "comies/jobs/ordering"

type Handler struct {
	statuses ordering.Jobs
}

func NewHandler(statuses ordering.Jobs) Handler {
	return Handler{
		statuses: statuses,
	}
}
