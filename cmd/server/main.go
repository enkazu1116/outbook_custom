package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	// 詳細実装前の簡易処理
	// Echoインスタンスの初期化
	e := echo.New()

	// ルートの設定
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// サーバーの起動
	// 失敗時はログに出力して終了
	e.Logger.Fatal(e.Start(":1322"))
}
