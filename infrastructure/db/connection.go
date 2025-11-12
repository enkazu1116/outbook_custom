package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"app/internal/domain/user/entity"
)

func NewConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// マイグレーション
	if err := db.AutoMigrate(&entity.User{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	return db
}
