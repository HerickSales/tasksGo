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
	Nome         string `json:"nome"`
	CriadorID    uint   `json:"criadorID"`
	ConcluinteID *uint  `json:"concluinteID,omitempty"`
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

	if r.ConcluinteID != nil {
		if err := db.First(&user, *r.ConcluinteID).Error; err != nil {
			return fmt.Errorf("user with id %d does not exist", *r.ConcluinteID)
		}
	}

	if !user.IsAdmin {
		return fmt.Errorf("user with id %d is not an admin", r.CriadorID)
	}

	// if r.ConcluinteID != &user.ID {
	// 	return fmt.Errorf("user with id %d is not allowed to end this task", *r.ConcluinteID)
	// } (MIGRAR ESSA PARTE PARA UPDATE TASK)

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
	// a implementar
}
