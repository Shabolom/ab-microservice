package di

import (
	"ab/internal/repository/pg-repo/experiment"
	"ab/internal/repository/pg-repo/group"
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
