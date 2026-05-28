package handler

type ExampleService interface{}

type HTTPHandlers struct {
	ExampleHandler *ExampleHandler
}

func NewHTTPHandlers(exampleService ExampleService) *HTTPHandlers {
	return &HTTPHandlers{
		ExampleHandler: NewExampleHandler(exampleService),
	}
}
