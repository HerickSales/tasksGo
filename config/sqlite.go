package config

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSQLite() (*gorm.DB, error) {
	dbPath := "./db/main.db"

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("sqlite opening error: %v", err)
	}
	return db, nil
}
