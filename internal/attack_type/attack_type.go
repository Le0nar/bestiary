package attacktype

type AttackType struct {
	Id    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Range int    `json:"range" db:"range"`
}
