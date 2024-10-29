package repository

import "github.com/jmoiron/sqlx"

const npcTable = "npc"
const attackTypeTable = "attack_type"
const enemyTable = "enemy"

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository  {
	return &Repository{db: db}
}

// func (r *Repository) CreateNpc () error {

// }