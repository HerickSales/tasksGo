package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model
	Nome              string
	IsAdmin           bool
	TarefasDesignadas []Tarefa
}

type UsuarioResponse struct {
	ID                uint      `json:"id"`
	CreateAt          time.Time `json:"createAt"`
	UpdateAt          time.Time `json:"updateAt"`
	DeletedAt         time.Time `json:"deletedAt,omitempty"`
	Nome              string    `json:"nome"`
	IsAdmin           bool      `json:"isAdmin"`
	TarefasDesignadas []Tarefa  `json:"tarefasDesignadas"`
}
