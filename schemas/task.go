package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Tarefa struct {
	gorm.Model
	Nome         string
	Concluida    bool
	CriadorID    uint     `gorm:"index"`
	Criador      Usuario  `gorm:"foreignKey:CriadorID"`
	ConcluinteID *uint    `gorm:"index"`
	Concluinte   *Usuario `gorm:"foreignKey:ConcluinteID"`
}

type TarefaResponse struct {
	ID         uint      `json:"id"`
	CreateAt   time.Time `json:"createAt"`
	UpdateAt   time.Time `json:"updateAt"`
	DeletedAt  time.Time `json:"deletedAt,omitempty"`
	Nome       string    `json:"nome"`
	Concluida  bool      `json:"concluida"`
	Criador    Usuario   `json:"criador"`
	Concluinte *Usuario  `json:"concluinte,omitempty"`
}
