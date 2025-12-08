package main

import (
	"app/infrastructure/di"
	"app/internal/application/interface/handler"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	// 詳細実装前の簡易処理
	// Echoインスタンスの初期化
	e := echo.New()

	// DI済みのAppオブジェクトの取得
	app := di.InitializeApp()

	// ハンドラの作成
	userHandler := handler.NewUserHandler(app.CreateUserUseCase)

	// ルーティング
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/users", userHandler.CreateUser)

	// サーバーの起動
	// 失敗時はログに出力して終了
	e.Logger.Fatal(e.Start(":1322"))
}
