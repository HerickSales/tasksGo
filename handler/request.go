package handler

import (
	"fmt"

	"github.com/elipe1/tasksGo-tarefa2/schemas"
)

type CreateUserRequest struct {
	Nome string `json:"nome"`
}

func (r *CreateUserRequest) Validate() error {
	if r.Nome == "" {
		return fmt.Errorf("nome is required")
	}
	return nil
}

type CreateTaskRequest struct {
	Nome    string          `json:"nome"`
	Criador schemas.Usuario `json:"criador"`
}

func (r *CreateTaskRequest) Validate() error {
	if r.Nome == "" {
		return fmt.Errorf("nome is required")
	}
	if r.Criador.Nome == "" {
		return fmt.Errorf("Criador.Nome is required")
	}
	return nil

}
