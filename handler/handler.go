package handler

import (
	"github.com/elipe1/tasksGo-tarefa2/config"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitializeHandler() {
	db = config.GetSQLite()
}
