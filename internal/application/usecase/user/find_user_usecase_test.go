package user

import (
	userdto "app/internal/application/dto/user"
	"app/internal/domain/user/entity"
	repo "app/internal/domain/user/repository"
	"context"
	"errors"
	"testing"
)

const searchRequiredMessage = "検索条件を1つ以上指定してください。"

func TestFindUserUsecase_FindUser(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	t.Run("validation error", func(t *testing.T) {
		t.Parallel()

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
