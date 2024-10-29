package handler

import (
	"net/http"

	"github.com/go-chi/chi"
)

type service interface{}

type Handler struct {
	service service
}

func NewHandler(s service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) InitRouter() http.Handler {
	router := chi.NewRouter()

	return router
}