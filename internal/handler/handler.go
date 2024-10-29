package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Le0nar/bestiary/internal/npc"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type service interface{
	CreateNpc(dto npc.CreateNpcDto) error
	GetNpcList() ([]npc.Npc, error)
}

type Handler struct {
	service service
}

func NewHandler(s service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) InitRouter() http.Handler {
	router := chi.NewRouter()

	router.Post("/npc", h.CreateNpc)
	router.Get("/npc", h.GetNpcList)

	return router
}

func (h *Handler) CreateNpc(w http.ResponseWriter, r *http.Request) {
	var dto npc.CreateNpcDto

	err:= json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.CreateNpc(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, "")
}

func (h *Handler) GetNpcList(w http.ResponseWriter, r *http.Request)  {
	npcList, err := h.service.GetNpcList();
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, npcList)
}
