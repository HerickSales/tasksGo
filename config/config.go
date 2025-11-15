package config

import (
	"fmt"

	"github.com/elipe1/tasksGo-tarefa2/schemas"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {
	var err error

	db, err = InitSQLite()
	if err != nil {
		return fmt.Errorf("error initialize sqlite: %v", err)
	}

	err = db.AutoMigrate(&schemas.Usuario{}, &schemas.Tarefa{})
	if err != nil {
		return fmt.Errorf("error migrating database: %v", err)
	}

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}
