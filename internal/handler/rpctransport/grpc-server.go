package rpctransport

import (
	customParams "ab/internal/handler/rpctransport/custom-params"
	"ab/internal/handler/rpctransport/experiments"
	"ab/internal/handler/rpctransport/layer"
	"ab/internal/handler/rpctransport/namespace"
)

type (
	experimentsHandler = *experiments.Handler
	namespaceHandler   = *namespace.Handler
	layerHandler       = *layer.Handler
	customParam        = *customParams.Handler
)

type Handlers struct {
	experimentsHandler
	namespaceHandler
	layerHandler
	customParam
}

func New(
	getExperimentsHandler experimentsHandler,
	namespaceHandler namespaceHandler,
	layerHandler layerHandler,
	customParam customParam,
) *Handlers {
	return &Handlers{
		experimentsHandler: getExperimentsHandler,
		namespaceHandler:   namespaceHandler,
		layerHandler:       layerHandler,
		customParam:        customParam,
	}
}
