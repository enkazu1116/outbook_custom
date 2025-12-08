package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	userdto "app/internal/application/dto/user"
	usecase "app/internal/application/usecase/user"
	"app/internal/domain/user/entity"
	"app/internal/domain/user/repository"
	"app/internal/domain/user/value_obj"
	testlogger "app/internal/test/logger"

	"github.com/labstack/echo/v4"
)

// testUserRepository はハンドラー経由で呼ばれる CreateUserUsecase 用のテストリポジトリです。
//
// ハンドラーのテストでは、「ユースケースの内部でどのような結果（成功/失敗）が起きるか」を
// 制御する目的で利用します。ExistsByEmail や CreateUser の戻り値を変えることで、
// HTTP レスポンスのステータスコードがどう変化するかを簡単に確認できます。
type testUserRepository struct {
	existsByEmailFn func(ctx context.Context, email string) (bool, error)
	createUserFn    func(ctx context.Context, u *entity.User) error
}

func (m *testUserRepository) CreateUser(ctx context.Context, u *entity.User) error {
	if m.createUserFn != nil {
		return m.createUserFn(ctx, u)
	}
	return nil
}

func (m *testUserRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	if m.existsByEmailFn != nil {
		return m.existsByEmailFn(ctx, email)
	}
	return false, nil
}

func (m *testUserRepository) FindByUser(context.Context, string, string, string) (*entity.User, error) {
	return nil, nil
}

func (m *testUserRepository) UpdateUser(context.Context, *entity.User) error {
	return nil
}

func (m *testUserRepository) DeleteUser(context.Context, string) error {
	return nil
}

var _ repository.UserRepository = (*testUserRepository)(nil)

// testPasswordHasher は CreateUserUsecase 用のテストハッシャーです。
// ハンドラーのテストではハッシュロジック自体は重要ではないため、
// ここでは常に成功するシンプルな実装を用意して、ハンドラーとユースケース間の連携に集中します。
type testPasswordHasher struct {
	hashFn func(password string) (string, error)
}

func (m *testPasswordHasher) Hash(password string) (string, error) {
	if m.hashFn != nil {
		return m.hashFn(password)
	}
	return "hashed-" + password, nil
}

func (m *testPasswordHasher) Compare(password, hash string) bool {
	return true
}

// TestUserHandler_CreateUser はユーザー作成ハンドラーの挙動をテストします。
//
// - Bind 失敗時に 400 を返す
// - ユースケースがエラーを返した場合に 500 を返す
// - 正常系で 201 を返す
//
// という典型的な Web ハンドラの振る舞いを TDD で固定しておくことで、
// 将来ハンドラー内部をリファクタリングしても、外から見た振る舞いが崩れていないかを素早く確認できます。
func TestUserHandler_CreateUser(t *testing.T) {
	t.Parallel()

	logger := testlogger.New(t)
	logger.Info(value_obj.UserUsecaseTestStartInfo.Message())
	defer logger.Info(value_obj.UserUsecaseTestSuccessInfo.Message())

	e := echo.New()

	t.Run("bind error returns 400", func(t *testing.T) {
		t.Parallel()

		logger := testlogger.New(t)
		logger.Info("UserHandler CreateUser Bind エラーケース開始")

		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(`invalid-json`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Bind エラーのケースではユースケースは呼ばれないため nil でも問題ありません。
		h := NewUserHandler((*usecase.CreateUserUsecase)(nil))

		if err := h.CreateUser(c); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("status code = %d, want %d", rec.Code, http.StatusBadRequest)
		}
	})

	t.Run("usecase error returns 500", func(t *testing.T) {
		t.Parallel()

		logger := testlogger.New(t)
		logger.Info("UserHandler CreateUser ユースケースエラーケース開始")

		body, _ := json.Marshal(userdto.CreateUserCommand{
			Name:     "Alice",
			Email:    "alice@example.com",
			Password: "Password1",
			Bio:      "hello",
		})

		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		repoMock := &testUserRepository{
			existsByEmailFn: func(_ context.Context, _ string) (bool, error) {
				// メールアドレス重複として扱わせ、ユースケース側でエラーを発生させる
				return true, nil
			},
		}
		hasherMock := &testPasswordHasher{}

		uc := usecase.NewCreateUserUsecase(repoMock, hasherMock)
		h := NewUserHandler(uc)

		if err := h.CreateUser(c); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if rec.Code != http.StatusInternalServerError {
			t.Fatalf("status code = %d, want %d", rec.Code, http.StatusInternalServerError)
		}
	})

	t.Run("success returns 201", func(t *testing.T) {
		t.Parallel()

		logger := testlogger.New(t)
		logger.Info("UserHandler CreateUser 正常系ケース開始")

		body, _ := json.Marshal(userdto.CreateUserCommand{
			Name:     "Alice",
			Email:    "alice@example.com",
			Password: "Password1",
			Bio:      "hello",
		})

		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		repoMock := &testUserRepository{
			existsByEmailFn: func(_ context.Context, _ string) (bool, error) {
				return false, nil
			},
			createUserFn: func(_ context.Context, _ *entity.User) error {
				return nil
			},
		}
		hasherMock := &testPasswordHasher{}

		uc := usecase.NewCreateUserUsecase(repoMock, hasherMock)
		h := NewUserHandler(uc)

		if err := h.CreateUser(c); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if rec.Code != http.StatusCreated {
			t.Fatalf("status code = %d, want %d", rec.Code, http.StatusCreated)
		}
	})
}


