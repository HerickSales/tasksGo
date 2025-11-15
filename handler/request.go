package handler

import (
	"fmt"

	"github.com/elipe1/tasksGo-tarefa2/schemas"
)

type CreateUserRequest struct {
	Nome    string `json:"nome"`
	Senha   string `json:"senha"`
	IsAdmin bool   `json:"isAdmin"`
}

type LoginRequest struct {
	Nome  string `json:"nome"`
	Senha string `json:"senha"`
}

func (r *CreateUserRequest) Validate() error {
	if r.Nome == "" {
		return fmt.Errorf("nome is required")
	}
	return nil
}

type CreateTaskRequest struct {
	Nome         string `json:"nome"`
	CriadorID    uint   `json:"criadorID"`
	ConcluinteID *uint  `json:"concluinteID"`
}

func (r *CreateTaskRequest) Validate() error {
	if r.Nome == "" {
		return fmt.Errorf("nome is required")
	}
	if r.CriadorID == 0 {
		return fmt.Errorf("CriadorID is required")
	}

	var criador schemas.Usuario
	if err := db.First(&criador, r.CriadorID).Error; err != nil {
		return fmt.Errorf("criador with id %d does not exist", r.CriadorID)
	}

	if !criador.IsAdmin {
		return fmt.Errorf("user with id %d is not an admin", r.CriadorID)
	}

	if r.ConcluinteID != nil {
		var concluinte schemas.Usuario
		if err := db.First(&concluinte, *r.ConcluinteID).Error; err != nil {
			return fmt.Errorf("user with id %d does not exist", *r.ConcluinteID)
		}
	}

	return nil
}

type UpdateUserRequest struct {
	Nome    string `json:"nome"`
	IsAdmin *bool  `json:"isAdmin"`
}

func (r *UpdateUserRequest) Validate() error {
	if r.Nome != "" || r.IsAdmin != nil {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}

type UpdateTaskRequest struct {
	Nome         string `json:"nome"`
	Concluida    *bool  `json:"concluida,omitempty"`
	ConcluinteID *uint  `json:"concluinteID,omitempty"`
}

func (r *UpdateTaskRequest) Validate() error {
	if r.Nome == "" && r.ConcluinteID == nil {
		return fmt.Errorf("at least one valid field must be provided")
	}

	if r.ConcluinteID != nil {
		var user schemas.Usuario
		if err := db.First(&user, *r.ConcluinteID).Error; err != nil {
			return fmt.Errorf("user with id %d does not exist", *r.ConcluinteID)
		}
	}

	// Ao editar uma tarefa, solicitar o id do usuário que está fazendo a edição e verificar se ele é admin

	return nil
}
