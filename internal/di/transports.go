package di

import "ab/internal/handler"

func (d *DI) GetHTTPHandlers() *handler.HTTPHandlers {
	return handler.NewHTTPHandlers(
		d.GetExampleService(),
	)
}
