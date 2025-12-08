package user

import (
	userdto "app/internal/application/dto/user"
	"app/internal/application/port"
	"app/internal/domain/user/entity"
	"app/internal/domain/user/value_obj"
	repo "app/internal/domain/user/repository"
	testlogger "app/internal/test/logger"
	"context"
	"errors"
	"testing"
)

// testPasswordHasher は PasswordHasher ポートを満たすテスト用の実装です。
// ユースケースが「ハッシュ化そのもののロジック」ではなく「ハッシュ化の呼び出し方」に
// 正しく責務を限定できているかを確認するために利用します。
type testPasswordHasher struct {
	hashFn    func(password string) (string, error)
	compareFn func(password, hash string) bool
}

func (m *testPasswordHasher) Hash(password string) (string, error) {
	if m.hashFn != nil {
		return m.hashFn(password)
	}
	return "", errors.New("hashFn not set")
}

func (m *testPasswordHasher) Compare(password, hash string) bool {
	if m.compareFn != nil {
		return m.compareFn(password, hash)
	}
	return false
}

// testCreateUserRepository は UserRepository を満たすテスト用実装です。
// 各テストケースで「重複チェック結果」や「CreateUser 時のエラー有無」を細かく制御できるようにし、
// ユースケースの振る舞いだけにフォーカスして検証する目的で使います。
type testCreateUserRepository struct {
	createUserFn    func(ctx context.Context, u *entity.User) error
	existsByEmailFn func(ctx context.Context, email string) (bool, error)
}

func (m *testCreateUserRepository) CreateUser(ctx context.Context, u *entity.User) error {
	if m.createUserFn != nil {
		return m.createUserFn(ctx, u)
	}
	return nil
}

func (m *testCreateUserRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	if m.existsByEmailFn != nil {
		return m.existsByEmailFn(ctx, email)
	}
	return false, nil
}

func (m *testCreateUserRepository) FindByUser(context.Context, string, string, string) (*entity.User, error) {
	return nil, errors.New("not implemented")
}

func (m *testCreateUserRepository) UpdateUser(context.Context, *entity.User) error {
	return errors.New("not implemented")
}

func (m *testCreateUserRepository) DeleteUser(context.Context, string) error {
	return errors.New("not implemented")
}

var _ repo.UserRepository = (*testCreateUserRepository)(nil)
var _ port.PasswordHasher = (*testPasswordHasher)(nil)

// TestCreateUserUsecase_CreateUser はユーザー作成ユースケースの振る舞いを一通り検証します。
// 入力バリデーション・メールアドレス重複・リポジトリエラー・ハッシュエラー・正常系など、
// ユースケースとして想定される代表的なシナリオを網羅します。
func TestCreateUserUsecase_CreateUser(t *testing.T) {
	t.Parallel()

	logger := testlogger.New(t)
	logger.Info(value_obj.UserUsecaseTestStartInfo.Message())
	defer logger.Info(value_obj.UserUsecaseTestSuccessInfo.Message())

	ctx := context.Background()

	t.Run("validation error", func(t *testing.T) {
		t.Parallel()

		logger := testlogger.New(t)
		logger.Info("CreateUserUsecase バリデーションエラーケース開始")

		uc := NewCreateUserUsecase(&testCreateUserRepository{}, &testPasswordHasher{})

		cmd := userdto.CreateUserCommand{}
		if err := uc.CreateUser(ctx, cmd); err == nil {
			t.Fatal("expected validation error, got nil")
		}
	})

	t.Run("email already exists", func(t *testing.T) {
		t.Parallel()

		logger := testlogger.New(t)
		logger.Info("CreateUserUsecase メールアドレス重複ケース開始")

		repoMock := &testCreateUserRepository{
			existsByEmailFn: func(_ context.Context, email string) (bool, error) {
				if email != "alice@example.com" {
					t.Fatalf("unexpected email: %s", email)
				}
				return true, nil
			},
		}
		hasherMock := &testPasswordHasher{
			hashFn: func(password string) (string, error) {
				return "hashed-" + password, nil
			},
		}

		uc := NewCreateUserUsecase(repoMock, hasherMock)

		cmd := userdto.CreateUserCommand{
			Name:     "Alice",
			Email:    "alice@example.com",
			Password: "Password1",
			Bio:      "hello",
		}

		if err := uc.CreateUser(ctx, cmd); err == nil {
			t.Fatal("expected email exists error, got nil")
		}
	})

	t.Run("existsByEmail error", func(t *testing.T) {
		t.Parallel()

		logger := testlogger.New(t)
		logger.Info("CreateUserUsecase ExistsByEmail エラーケース開始")

		expectedErr := errors.New("db error")
		repoMock := &testCreateUserRepository{
			existsByEmailFn: func(_ context.Context, _ string) (bool, error) {
				return false, expectedErr
			},
		}

		uc := NewCreateUserUsecase(repoMock, &testPasswordHasher{
			hashFn: func(password string) (string, error) {
				return "hashed-" + password, nil
			},
		})

		cmd := userdto.CreateUserCommand{
			Name:     "Alice",
			Email:    "alice@example.com",
			Password: "Password1",
			Bio:      "hello",
		}

		err := uc.CreateUser(ctx, cmd)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, expectedErr) {
			t.Fatalf("expected wrapped error %v, got %v", expectedErr, err)
		}
	})

	t.Run("hash error", func(t *testing.T) {
		t.Parallel()

		logger := testlogger.New(t)
		logger.Info("CreateUserUsecase ハッシュエラーケース開始")

		expectedErr := errors.New("hash error")
		repoMock := &testCreateUserRepository{
			existsByEmailFn: func(_ context.Context, _ string) (bool, error) {
				return false, nil
			},
		}
		hasherMock := &testPasswordHasher{
			hashFn: func(_ string) (string, error) {
				return "", expectedErr
			},
		}

		uc := NewCreateUserUsecase(repoMock, hasherMock)

		cmd := userdto.CreateUserCommand{
			Name:     "Alice",
			Email:    "alice@example.com",
			Password: "Password1",
			Bio:      "hello",
		}

		err := uc.CreateUser(ctx, cmd)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, expectedErr) {
			t.Fatalf("expected wrapped error %v, got %v", expectedErr, err)
		}
	})

	t.Run("create error", func(t *testing.T) {
		t.Parallel()

		logger := testlogger.New(t)
		logger.Info("CreateUserUsecase CreateUser エラーケース開始")

		expectedErr := errors.New("create error")
		repoMock := &testCreateUserRepository{
			existsByEmailFn: func(_ context.Context, _ string) (bool, error) {
				return false, nil
			},
			createUserFn: func(_ context.Context, _ *entity.User) error {
				return expectedErr
			},
		}
		hasherMock := &testPasswordHasher{
			hashFn: func(password string) (string, error) {
				return "hashed-" + password, nil
			},
		}

		uc := NewCreateUserUsecase(repoMock, hasherMock)

		cmd := userdto.CreateUserCommand{
			Name:     "Alice",
			Email:    "alice@example.com",
			Password: "Password1",
			Bio:      "hello",
		}

		err := uc.CreateUser(ctx, cmd)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, expectedErr) {
			t.Fatalf("expected wrapped error %v, got %v", expectedErr, err)
		}
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		logger := testlogger.New(t)
		logger.Info("CreateUserUsecase 正常系ケース開始")

		var created *entity.User

		repoMock := &testCreateUserRepository{
			existsByEmailFn: func(_ context.Context, email string) (bool, error) {
				if email != "alice@example.com" {
					t.Fatalf("unexpected email: %s", email)
				}
				return false, nil
			},
			createUserFn: func(_ context.Context, u *entity.User) error {
				created = u
				return nil
			},
		}

		hasherMock := &testPasswordHasher{
			hashFn: func(password string) (string, error) {
				if password != "Password1" {
					t.Fatalf("unexpected password: %s", password)
				}
				return "hashed-Password1", nil
			},
		}

		uc := NewCreateUserUsecase(repoMock, hasherMock)

		cmd := userdto.CreateUserCommand{
			Name:     "Alice",
			Email:    "alice@example.com",
			Password: "Password1",
			Bio:      "hello",
		}

		if err := uc.CreateUser(ctx, cmd); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if created == nil {
			t.Fatal("expected user to be created, got nil")
		}
		if created.Email != "alice@example.com" {
			t.Errorf("created.Email = %s, want %s", created.Email, "alice@example.com")
		}
		if created.Password != "hashed-Password1" {
			t.Errorf("created.Password = %s, want %s", created.Password, "hashed-Password1")
		}
	})
}


