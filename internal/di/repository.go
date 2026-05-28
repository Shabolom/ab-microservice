package di

import "context"

type ExampleRepository interface {
	Ping(ctx context.Context) error
}

func (d *DI) GetExampleRepository() ExampleRepository {
	// TODO: replace with real repository constructor
	return nil
}
