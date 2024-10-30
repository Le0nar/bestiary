package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Le0nar/bestiary/internal/creature"
	"github.com/Le0nar/bestiary/internal/enemy"
	"github.com/Le0nar/bestiary/internal/npc"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type service interface{
	CreateNpc(dto npc.CreateNpcDto) error
	GetNpcList() ([]npc.Npc, error)

	CreateEnemy(dto enemy.CreateEnemyDto) error
	GetEnemyList() ([]enemy.EnemyDto, error)

	GetCreatureList()([]creature.Creature, error)
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

	router.Post("/enemy", h.CreateEnemy)
	router.Get("/enemy", h.GetEnemyList)

	router.Get("/creature", h.GetCreatureList)

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

func (h *Handler) CreateEnemy(w http.ResponseWriter, r *http.Request)  {
	var dto enemy.CreateEnemyDto

	err:= json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.CreateEnemy(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, "")
}

func (h *Handler) GetEnemyList(w http.ResponseWriter, r *http.Request)  {
	enemyList, err := h.service.GetEnemyList();
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, enemyList)
}

func (h *Handler) GetCreatureList(w http.ResponseWriter, r *http.Request)  {
	creature, err := h.service.GetCreatureList();
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, creature)
}
