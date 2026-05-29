package di

import (
	"ab/internal/handler/rpctransport"
	"ab/internal/handler/rpctransport/experiments"
)

func (d *DI) GetGRPCHandlers() *rpctransport.Handlers {
	return rpctransport.New(d.GetExperimentHandler())
}

func (d *DI) GetExperimentHandler() *experiments.Handler {
	return experiments.New(d.GetExampleService())
}
