package npc

type Npc struct {
	Id    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Description  string `json:"description" db:"description"`
	HP int    `json:"hp" db:"hp"`
}

type CreateNpcDto struct {
	Name  string `json:"name" db:"name"`
	Description  string `json:"description" db:"description"`
	HP int    `json:"hp" db:"hp"`
}
