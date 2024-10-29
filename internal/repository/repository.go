package repository

import (
	"fmt"

	"github.com/Le0nar/bestiary/internal/npc"
	"github.com/jmoiron/sqlx"
)

const npcTable = "npc"
const attackTypeTable = "attack_type"
const enemyTable = "enemy"

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository  {
	return &Repository{db: db}
}

func (r *Repository) CreateNpc(dto npc.CreateNpcDto) error {
	query := fmt.Sprintf("INSERT INTO %s (name, description, hp) values ($1, $2, $3)", npcTable)

	_, err := r.db.Exec(query, dto.Name, dto.Description, dto.HP)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetNpcList() ([]npc.Npc, error) {
	var npcList []npc.Npc

	query := fmt.Sprintf("SELECT * FROM %s", npcTable)

	err := r.db.Select(
		&npcList,
		query,
	)

	return npcList, err
}
