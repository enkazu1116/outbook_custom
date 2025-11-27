package user

import (
	userdto "app/internal/application/dto/user"
	"app/internal/domain/user/entity"
	repo "app/internal/domain/user/repository"
	"app/internal/domain/user/value_obj"
	testlogger "app/internal/test/logger"
	"context"
	"errors"
	"testing"
)

const searchRequiredMessage = "検索条件を1つ以上指定してください。"

// TestFindUserUsecase_FindUser はユーザー検索ユースケースの動作を検証します。
//
// バリデーションエラー・リポジトリ成功・リポジトリエラーの各シナリオを通じて、
// 「どのような入力のときに、ユースケースがどのようにリポジトリを呼び出すべきか」を
// 一目で思い出せるようにすることを意図しています。
func TestFindUserUsecase_FindUser(t *testing.T) {
	t.Parallel()

	logger := testlogger.New(t)
	logger.Info(value_obj.UserUsecaseTestStartInfo.Message())
	defer logger.Info(value_obj.UserUsecaseTestSuccessInfo.Message())

	ctx := context.Background()

	t.Run("validation error", func(t *testing.T) {
		t.Parallel()

		logger := testlogger.New(t)
		logger.Info("FindUserUsecase バリデーションエラーケース開始")

		uc := NewFindUserUsecase(&testUserRepository{})

		_, err := uc.FindUser(ctx, userdto.FindUserQuery{})
		if err == nil {
			t.Fatalf("expected validation error, got nil")
		}
		if err.Error() != searchRequiredMessage {
			t.Fatalf("unexpected error message: %v", err)
		}
	})

	t.Run("repository success", func(t *testing.T) {
		t.Parallel()

		logger := testlogger.New(t)
		logger.Info("FindUserUsecase リポジトリ成功ケース開始")

		expected := &entity.User{ID: "user-1", Name: "Alice", Email: "alice@example.com"}
		mock := &testUserRepository{
			findByUserFn: func(_ context.Context, id, name, email string) (*entity.User, error) {
				if id != "user-1" || name != "" || email != "" {
					t.Fatalf("unexpected parameters: id=%s name=%s email=%s", id, name, email)
				}
				return expected, nil
			},
		}

		uc := NewFindUserUsecase(mock)

		user, err := uc.FindUser(ctx, userdto.FindUserQuery{ID: "user-1"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if user != expected {
			t.Fatalf("expected user pointer %p, got %p", expected, user)
		}
	})

	t.Run("repository error", func(t *testing.T) {
		t.Parallel()

		logger := testlogger.New(t)
		logger.Info("FindUserUsecase リポジトリエラーケース開始")

		expectedErr := errors.New("db error")
		mock := &testUserRepository{
			findByUserFn: func(_ context.Context, _, _, _ string) (*entity.User, error) {
				return nil, expectedErr
			},
		}

		uc := NewFindUserUsecase(mock)

		_, err := uc.FindUser(ctx, userdto.FindUserQuery{Email: "alice@example.com"})
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, expectedErr) {
			t.Fatalf("expected wrapped error %v, got %v", expectedErr, err)
		}
	})
}

type testUserRepository struct {
	findByUserFn func(ctx context.Context, id, name, email string) (*entity.User, error)
}

func (m *testUserRepository) CreateUser(context.Context, *entity.User) error {
	return errors.New("not implemented")
}

func (m *testUserRepository) ExistsByEmail(context.Context, string) (bool, error) {
	return false, errors.New("not implemented")
}

func (m *testUserRepository) FindByUser(ctx context.Context, id string, name string, email string) (*entity.User, error) {
	if m.findByUserFn != nil {
		return m.findByUserFn(ctx, id, name, email)
	}
	return nil, errors.New("findByUserFn not set")
}

func (m *testUserRepository) UpdateUser(context.Context, *entity.User) error {
	return errors.New("not implemented")
}

func (m *testUserRepository) DeleteUser(context.Context, string) error {
	return errors.New("not implemented")
}

var _ repo.UserRepository = (*testUserRepository)(nil)
