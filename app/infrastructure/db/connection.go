package db

import (
	"database/sql"
	"os"

	"app/infrastructure/logger"

	libsql "github.com/tursodatabase/libsql-client-go/libsql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"app/internal/domain/user/entity"
)

func NewConnection() *gorm.DB {
	tursoURL := os.Getenv("TURSO_DATABASE_URL")
	tursoAuthToken := os.Getenv("TURSO_AUTH_TOKEN")

	if tursoURL == "" || tursoAuthToken == "" {
		logger.FatalJp("環境変数 TURSO_DATABASE_URL と TURSO_AUTH_TOKEN は必須です")
	}

	// Tursoコネクターの作成
	connector, err := libsql.NewConnector(tursoURL, libsql.WithAuthToken(tursoAuthToken))
	if err != nil {
		logger.FatalJp("Turso への接続コネクター作成に失敗しました: %v", err)
	}

	// database/sqlでTursoに接続
	sqlDB := sql.OpenDB(connector)

	// 接続テスト
	if err := sqlDB.Ping(); err != nil {
		logger.FatalJp("データベースへの接続確認に失敗しました: %v", err)
	}

	// GORMでTursoを使用
	db, err := gorm.Open(sqlite.Dialector{Conn: sqlDB}, &gorm.Config{})
	if err != nil {
		logger.FatalJp("データベース接続の初期化に失敗しました: %v", err)
	}

	// マイグレーション
	if err := db.AutoMigrate(&entity.User{}); err != nil {
		logger.FatalJp("ユーザーテーブルのマイグレーションに失敗しました: %v", err)
	}

	return db
}
