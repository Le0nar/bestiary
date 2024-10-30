package enemy

type Enemy struct {
	Id    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Description  string `json:"description" db:"description"`
	HP int    `json:"hp" db:"hp"`
	AttackType int    `json:"attackType" db:"attack_type"`
	Damage int    `json:"damage" db:"damage"`
	Haste int    `json:"haste" db:"haste"`
}

type EnemyDto struct {
	Id    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Description  string `json:"description" db:"description"`
	HP int    `json:"hp" db:"hp"`
	AttackType string    `json:"attackType" db:"attack_type"`
	Damage int    `json:"damage" db:"damage"`
	Haste int    `json:"haste" db:"haste"`
}

type CreateEnemyDto struct {
	Name  string `json:"name" db:"name"`
	Description  string `json:"description" db:"description"`
	HP int    `json:"hp" db:"hp"`
	AttackType string    `json:"attackType" db:"attack_type"`
	Damage int    `json:"damage" db:"damage"`
	Haste int    `json:"haste" db:"haste"`
}
