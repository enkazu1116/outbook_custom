package services

import "app/internal/domain/user/repository"

// UserRepository DI
type UserService struct {
	userRepository repository.UserRepository
}

// UserSeviceコンストラクタ
func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}
