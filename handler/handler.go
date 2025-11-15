package handler

import (
	"github.com/elipe1/tasksGo-tarefa2/config"
	"github.com/elipe1/tasksGo-tarefa2/schemas"
	"gorm.io/gorm"
)

var (
	db       *gorm.DB
	usuarios []schemas.Usuario
)

func InitializeHandler() {
	db = config.GetSQLite()
}
