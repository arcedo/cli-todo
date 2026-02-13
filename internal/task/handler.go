package task

type Handler struct {
	r Repository
}

func NewHandler(r Repository) *Handler {
	return &Handler{r}
}

func (h *Handler) Run(args []string) error {
	return nil
}
