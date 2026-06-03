package rpctransport

import (
	"ab/internal/handler/rpctransport/experiments"
	"ab/internal/handler/rpctransport/layer"
	"ab/internal/handler/rpctransport/namespace"
)

type (
	experimentsHandler = *experiments.Handler
	nameSpaceHandler   = *namespace.Handler
	layerHandler       = *layer.Handler
)

type Handlers struct {
	experimentsHandler
	nameSpaceHandler
	layerHandler
}

func New(
	getExperimentsHandler experimentsHandler,
	nameSpaceHandler nameSpaceHandler,
	layerHandler layerHandler,
) *Handlers {
	return &Handlers{
		experimentsHandler: getExperimentsHandler,
		nameSpaceHandler:   nameSpaceHandler,
		layerHandler:       layerHandler,
	}
}
