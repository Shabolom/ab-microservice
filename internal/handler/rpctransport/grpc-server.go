package rpctransport

import (
	customParams "ab/internal/handler/rpctransport/custom-params"
	"ab/internal/handler/rpctransport/experiments"
	"ab/internal/handler/rpctransport/layer"
	"ab/internal/handler/rpctransport/namespace"
)

type (
	experimentsHandler = *experiments.Handler
	nameSpaceHandler   = *namespace.Handler
	layerHandler       = *layer.Handler
	customParam        = *customParams.Handler
)

type Handlers struct {
	experimentsHandler
	nameSpaceHandler
	layerHandler
	customParam
}

func New(
	getExperimentsHandler experimentsHandler,
	nameSpaceHandler nameSpaceHandler,
	layerHandler layerHandler,
	customParam customParam,
) *Handlers {
	return &Handlers{
		experimentsHandler: getExperimentsHandler,
		nameSpaceHandler:   nameSpaceHandler,
		layerHandler:       layerHandler,
		customParam:        customParam,
	}
}
