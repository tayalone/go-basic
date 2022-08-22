package realworld

import "fmt"

type Handler struct {
	repo Repository
}

func NewHandler(repo Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) Do(id uint) string {
	lang := h.repo.QueryLang(id)
	return fmt.Sprintf("%s language", lang.Name)
}
