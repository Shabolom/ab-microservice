package di

import (
	"ab/internal/repository/pg-repo/experiment"
	"ab/internal/repository/pg-repo/group"
	"ab/internal/repository/pg-repo/layer"
	nameSpace "ab/internal/repository/pg-repo/name-space"
)

func (d *DI) GetExperimentPgRepo() *experiment.Storage {
	return experiment.New(d.GetPgDatabase())
}

func (d *DI) GetGroupPgRepo() *group.Storage {
	return group.New(d.GetPgDatabase())
}

func (d *DI) GetNamespacePgRepo() *nameSpace.Storage {
	return nameSpace.New(d.GetPgDatabase())
}

func (d *DI) GetLayerPgRepo() *layer.Storage {
	return layer.New(d.GetPgDatabase())
}
