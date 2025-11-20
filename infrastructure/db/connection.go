package db

import (
	"database/sql"
	"log"
	"os"

	libsql "github.com/tursodatabase/libsql-client-go/libsql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"app/internal/domain/user/entity"
)

func NewConnection() *gorm.DB {
	tursoURL := os.Getenv("TURSO_DATABASE_URL")
	tursoAuthToken := os.Getenv("TURSO_AUTH_TOKEN")

	if tursoURL == "" || tursoAuthToken == "" {
		log.Fatal("TURSO_DATABASE_URL and TURSO_AUTH_TOKEN environment variables are required")
	}

	// Tursoコネクターの作成
	connector, err := libsql.NewConnector(tursoURL, libsql.WithAuthToken(tursoAuthToken))
	if err != nil {
		log.Fatalf("failed to create Turso connector: %v", err)
	}

	// database/sqlでTursoに接続
	sqlDB := sql.OpenDB(connector)

	// 接続テスト
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	// GORMでTursoを使用
	db, err := gorm.Open(sqlite.Dialector{Conn: sqlDB}, &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// マイグレーション
	if err := db.AutoMigrate(&entity.User{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	return db
}
