package di

import (
	"ab/internal/service/experiments"
	"ab/internal/service/layer"
	"ab/internal/service/namespace"
)

func (d *DI) GetExperimentService() *experiments.Service {
	return experiments.New(
		d.GetGroupPgRepo(),
		d.GetExperimentPgRepo(),
		d.GetNamespacePgRepo(),
		d.GetInMemoryCache(),
		d.GetLayerPgRepo(),
		d.GetKafka(),
		d.Logger(),
	)
}

func (d *DI) GetLayerService() *layer.Service {
	return layer.New(d.Logger(), d.GetLayerPgRepo())
}

func (d *DI) GetNamespaceService() *namespace.Service {
	return namespace.New(d.Logger(), d.GetNamespacePgRepo())
}
