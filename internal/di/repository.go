package di

import (
	customParams "ab/internal/repository/pg-repo/custom-params"
	"ab/internal/repository/pg-repo/experiment"
	featureToggles "ab/internal/repository/pg-repo/feature-toggles"
	"ab/internal/repository/pg-repo/group"
	"ab/internal/repository/pg-repo/layer"
	"ab/internal/repository/pg-repo/namespace"
)

func (d *DI) GetExperimentPgRepo() *experiment.Storage {
	return experiment.New(d.GetPgDatabase())
}

func (d *DI) GetGroupPgRepo() *group.Storage {
	return group.New(d.GetPgDatabase())
}

func (d *DI) GetNamespacePgRepo() *namespace.Storage {
	return namespace.New(d.GetPgDatabase())
}

func (d *DI) GetLayerPgRepo() *layer.Storage {
	return layer.New(d.GetPgDatabase())
}

func (d *DI) GetCustomParamsPgRepo() *customParams.Storage {
	return customParams.New(d.GetPgDatabase())
}

func (d *DI) GetFeatureTogglePgRepo() *featureToggles.Storage {
	return featureToggles.New(d.GetPgDatabase())
}
