package handler

import (
	"net/http"

	"ab/internal/render"

	"github.com/labstack/echo/v4"
)

type ExampleHandler struct {
	service ExampleService
}

func NewExampleHandler(service ExampleService) *ExampleHandler {
	return &ExampleHandler{service: service}
}

func (h *ExampleHandler) Health(c echo.Context) error {
	return render.JSON(c, http.StatusOK, map[string]any{
		"status": "ok",
	})
}
