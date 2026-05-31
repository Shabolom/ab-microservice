package layer

type layerService interface {
}
type Handler struct {
	layerService layerService
}

func New(layerService layerService) *Handler {
	return &Handler{
		layerService: layerService,
	}
}
