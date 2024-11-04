package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Le0nar/bestiary/internal/creature"
	"github.com/Le0nar/bestiary/internal/enemy"
	"github.com/Le0nar/bestiary/internal/npc"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	httpSwagger "github.com/swaggo/http-swagger"
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

func (h *Handler) InitRouter(port int) http.Handler {
	router := chi.NewRouter()

	swaggerUrl := fmt.Sprintf("http://localhost:%d/swagger/doc.json", port)
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(swaggerUrl), //The url pointing to API definition
	))

	router.Post("/npc", h.CreateNpc)
	router.Get("/npc", h.GetNpcList)

	router.Post("/enemy", h.CreateEnemy)
	router.Get("/enemy", h.GetEnemyList)

	router.Get("/creature", h.GetCreatureList)

	return router
}

// @Summary Create npc
// @Description create npc
// @ID create-npc
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400,404
// @Failure 500
// @Failure default
// @Router /npc [post]
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

// @Summary Get npc list
// @Description Get npc list
// @ID get-npc-list
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400,404
// @Failure 500
// @Failure default
// @Router /npc [get]
func (h *Handler) GetNpcList(w http.ResponseWriter, r *http.Request)  {
	npcList, err := h.service.GetNpcList();
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, npcList)
}

// @Summary Create enemy
// @Description create enemy
// @ID create-enemy
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400,404
// @Failure 500
// @Failure default
// @Router /enemy [post]
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

// @Summary Get enemy list
// @Description Get enemy list
// @ID get-enemy-list
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400,404
// @Failure 500
// @Failure default
// @Router /enemy [get]
func (h *Handler) GetEnemyList(w http.ResponseWriter, r *http.Request)  {
	enemyList, err := h.service.GetEnemyList();
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, enemyList)
}

// @Summary Get creature list
// @Description Get creature list
// @ID get-creature-list
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400,404
// @Failure 500
// @Failure default
// @Router /creature [get]
func (h *Handler) GetCreatureList(w http.ResponseWriter, r *http.Request)  {
	creature, err := h.service.GetCreatureList();
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, creature)
}
