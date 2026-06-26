package rpctransport

import (
	customParams "ab/internal/handler/rpctransport/custom-params"
	"ab/internal/handler/rpctransport/experiments"
	featureToggles "ab/internal/handler/rpctransport/feature-toggles"
	"ab/internal/handler/rpctransport/layer"
	"ab/internal/handler/rpctransport/namespace"
)

type (
	experimentsHandler = *experiments.Handler
	namespaceHandler   = *namespace.Handler
	layerHandler       = *layer.Handler
	customParam        = *customParams.Handler
	featureToggle      = *featureToggles.Handler
)

type Handlers struct {
	experimentsHandler
	namespaceHandler
	layerHandler
	customParam
	featureToggle
}

func New(
	getExperimentsHandler experimentsHandler,
	namespaceHandler namespaceHandler,
	layerHandler layerHandler,
	customParam customParam,
	featureToggle featureToggle,
) *Handlers {
	return &Handlers{
		experimentsHandler: getExperimentsHandler,
		namespaceHandler:   namespaceHandler,
		layerHandler:       layerHandler,
		customParam:        customParam,
		featureToggle:      featureToggle,
	}
}
