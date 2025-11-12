package user

import (
	userdto "app/internal/application/dto/user"
	"app/internal/domain/user/entity"
	"app/internal/domain/user/repository"
	"app/internal/domain/user/services"
	"context"
	"fmt"
)

// FindUserUsecase はユーザー検索ユースケースを表します。
type FindUserUsecase struct {
	userRepository repository.UserRepository
}

// NewFindUserUsecase は FindUserUsecase のコンストラクタです。
func NewFindUserUsecase(userRepository repository.UserRepository) *FindUserUsecase {
	return &FindUserUsecase{userRepository: userRepository}
}

// FindUser は指定した条件でユーザーを検索します。
func (uc *FindUserUsecase) FindUser(ctx context.Context, query userdto.FindUserQuery) (*entity.User, error) {

	if err := services.FindUserValidation(ctx, query.ID, query.Name, query.Email); err != nil {
		return nil, err
	}

	user, err := uc.userRepository.FindByUser(ctx, query.ID, query.Name, query.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	return user, nil
}
