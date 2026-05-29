package rpctransport

import "ab/internal/handler/rpctransport/experiments"

type (
	getExperimentsHandler = *experiments.Handler
)

type Handlers struct {
	getExperimentsHandler
}

func New(getExperimentsHandler getExperimentsHandler) *Handlers {
	return &Handlers{
		getExperimentsHandler: getExperimentsHandler,
	}
}
