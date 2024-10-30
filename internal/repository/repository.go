package repository

import (
	"fmt"

	"github.com/Le0nar/bestiary/internal/creature"
	"github.com/Le0nar/bestiary/internal/enemy"
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

// Current transacation is overhead
func (r *Repository) CreateEnemy(dto enemy.CreateEnemyDto) error  {
	// 1) get index for attack type
	var attackTypeIndex int
	query := fmt.Sprintf("SELECT id FROM %s where name = '%s'", attackTypeTable, dto.AttackType)
	err := r.db.Get(&attackTypeIndex, query)

	if err != nil {
		return err
	}

	// 2) start transaction
	tx, err := r.db.Begin()

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	if err != nil {
		return err
	}

	// 3) set data to enemy
	query = fmt.Sprintf(
		"INSERT INTO %s (name, description, hp, attack_type, damage, haste) values ($1, $2, $3, $4, $5, $6)",
		 enemyTable,
	)

	_, err = tx.Exec(query, dto.Name, dto.Description, dto.HP, attackTypeIndex, dto.Damage, dto.Haste)
	if err != nil {
		return err
	}

	// 4) commit transaction
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetEnemyList() ([]enemy.EnemyDto, error) {
	var enemyList []enemy.Enemy

	query := fmt.Sprintf("SELECT * FROM %s", enemyTable)

	err := r.db.Select(
		&enemyList,
		query,
	)

	var enemyDtoList []enemy.EnemyDto

	for _, v  := range enemyList {
		var attackTypeName string
		query := fmt.Sprintf("SELECT name FROM %s where id = %d", attackTypeTable, v.AttackType)
		err := r.db.Get(&attackTypeName, query)

		if err != nil {
			return enemyDtoList, err
		}

		var enemyDto enemy.EnemyDto
		enemyDto.AttackType = attackTypeName
		enemyDto.Damage = v.Damage
		enemyDto.Description = v.Description
		enemyDto.HP = v.HP
		enemyDto.Haste = v.Haste
		enemyDto.Id = v.Id
		enemyDto.Name = v.Name

		enemyDtoList = append(enemyDtoList, enemyDto)

	}

	return enemyDtoList, err
}

func (r *Repository) GetCreatureList() ([]creature.Creature, error)  {
	var creatureList []creature.Creature

	query := fmt.Sprintf(
		"SELECT name, description FROM %s 		UNION ALL		SELECT name, description from %s		ORDER BY name", 
		npcTable,
		enemyTable,
	)

	err := r.db.Select(
		&creatureList,
		query,
	)

	return creatureList, err
}
