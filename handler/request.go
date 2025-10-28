package handler

import (
	"fmt"

	"github.com/elipe1/tasksGo-tarefa2/schemas"
)

type CreateUserRequest struct {
	Nome    string `json:"nome"`
	IsAdmin bool   `json:"isAdmin"`
}

func (r *CreateUserRequest) Validate() error {
	if r.Nome == "" {
		return fmt.Errorf("nome is required")
	}
	return nil
}

type CreateTaskRequest struct {
	Nome      string `json:"nome"`
	CriadorID uint   `json:"criadorID"`
}

func (r *CreateTaskRequest) Validate() error {
	if r.Nome == "" {
		return fmt.Errorf("nome is required")
	}
	if r.CriadorID == 0 {
		return fmt.Errorf("CriadorID is required")
	}

	var user schemas.Usuario
	if err := db.First(&user, r.CriadorID).Error; err != nil {
		return fmt.Errorf("user with id %d does not exist", r.CriadorID)
	}

	return nil
}
