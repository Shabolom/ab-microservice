package di

import "context"

type ExampleService interface {
	Health(ctx context.Context) error
}

func (d *DI) GetExampleService() ExampleService {
	// TODO: replace with real service constructor
	return nil
}
