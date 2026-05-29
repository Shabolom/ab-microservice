package di

import "ab/internal/repository/pg-repo/experiment"

func (d *DI) GetExperimentPgRepo() *experiment.Storage {
	return experiment.New(d.GetPgDatabase())
}

func (d *DI) GetGroupPgRepo() *experiment.Storage {
	return experiment.New(d.GetPgDatabase())
}
