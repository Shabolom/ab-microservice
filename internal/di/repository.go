package di

import "ab/internal/repository/pg-repo/experement"

func (d *DI) GetExperimentPgRepo() *experement.Storage {
	return experement.New(d.GetPgDatabase())
}

func (d *DI) GetGroupPgRepo() *experement.Storage {
	return experement.New(d.GetPgDatabase())
}
