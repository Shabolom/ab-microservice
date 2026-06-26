package di

import (
	"ab/internal/handler/rpctransport"
	"ab/internal/handler/rpctransport/custom-params"
	"ab/internal/handler/rpctransport/experiments"
	featureToggles "ab/internal/handler/rpctransport/feature-toggles"
	"ab/internal/handler/rpctransport/layer"
	"ab/internal/handler/rpctransport/namespace"
)

func (d *DI) GetGRPCHandlers() *rpctransport.Handlers {
	return rpctransport.New(
		d.GetExperimentHandler(),
		d.GetNamespaceHandler(),
		d.GetLayerHandler(),
		d.GetCustomParamsHandler(),
		d.GetFeatureTogglesHandler(),
	)
}

func (d *DI) GetNamespaceHandler() *namespace.Handler {
	return namespace.New(d.GetNamespaceService())
}

func (d *DI) GetLayerHandler() *layer.Handler {
	return layer.New(d.GetLayerService())
}

func (d *DI) GetExperimentHandler() *experiments.Handler {
	return experiments.New(d.GetExperimentService())
}

func (d *DI) GetCustomParamsHandler() *customParams.Handler {
	return customParams.New(d.GetCustomParamsService())
}

func (d *DI) GetFeatureTogglesHandler() *featureToggles.Handler {
	return featureToggles.New(d.GetFeatureTogglesService())
}
