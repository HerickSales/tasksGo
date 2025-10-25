package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Tarefa struct {
	gorm.Model
	Nome       string
	Concluida  bool
	Criador    Usuario
	Concluinte Usuario
}

type TarefaResponse struct {
	ID         uint      `json:"id"`
	CreateAt   time.Time `json:"createAt"`
	UpdateAt   time.Time `json:"updateAt"`
	DeletedAt  time.Time `json:"deletedAt,omitempty"`
	Nome       string    `json:"nome"`
	Concluida  bool      `json:"concluida"`
	Criador    Usuario   `json:"criador"`
	Concluinte Usuario   `json:"concluinte"`
}
