package namespace

type namespaceService interface {
}
type Handler struct {
	namespaceService namespaceService
}

func New(namespaceService namespaceService) *Handler {
	return &Handler{
		namespaceService: namespaceService,
	}
}
