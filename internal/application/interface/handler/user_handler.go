package handler

import (
	"app/internal/application/dto/user"
	usecase "app/internal/application/usecase/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

// UserHandler構造体
type UserHandler struct {
	usecase *usecase.CreateUserUsecase
}

// UserHandlerコンストラクタ
func NewUserHandler(uc *usecase.CreateUserUsecase) *UserHandler {
	return &UserHandler{usecase: uc}
}

// ユーザー作成ハンドラー
// 引数: echo.Context
// 返り値: エラー
// レシーバー: UserHandlerオブジェクト
func (h *UserHandler) CreateUser(c echo.Context) error {

	// DTOオブジェクトの作成
	var cmd user.CreateUserCommand

	// エラーハンドリング
	// Echoがバインド失敗時、エラーコードを返す
	if err := c.Bind(&cmd); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	// 実行 + エラーハンドリング
	// Usecaseが実行失敗時、エラーコードを返す
	if err := h.usecase.CreateUser(c.Request().Context(), cmd); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// NoContentで201 Createdを返す
	return c.NoContent(http.StatusCreated)
}
