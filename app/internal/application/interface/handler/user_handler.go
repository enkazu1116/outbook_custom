package handler

import (
	"app/internal/application/dto/user"
	usecase "app/internal/application/usecase/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

// UserHandler は HTTP レイヤからユーザー関連のユースケースを呼び出すためのハンドラです。
//
// この構造体自体は Echo の詳細（Context など）とアプリケーションユースケースの橋渡し役を担い、
// 具体的なビジネスロジックは CreateUserUsecase に委譲します。
type UserHandler struct {
	usecase *usecase.CreateUserUsecase
}

// NewUserHandler は UserHandler のコンストラクタです。
// ルーティング設定時にユースケースを注入して利用します。
func NewUserHandler(uc *usecase.CreateUserUsecase) *UserHandler {
	return &UserHandler{usecase: uc}
}

// CreateUser は HTTP 経由の「ユーザー作成リクエスト」を受け付けるハンドラです。
//
// 主な処理の流れは次の通りです。
//  1. リクエストボディ(JSON)を CreateUserCommand DTO にバインド
//  2. バインドに失敗した場合は 400 Bad Request を返却
//  3. ユースケース CreateUserUsecase.CreateUser を呼び出し
//  4. ユースケース側でエラーが発生した場合は 500 Internal Server Error を返却
//  5. 正常に作成できた場合は 201 Created（ボディ無し）を返却
//
// ここでは「リクエスト/レスポンスの形式」と「HTTP ステータスコードの決定」のみを担当し、
// 具体的なバリデーションやビジネスルールはユースケース・ドメイン層に任せています。
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
