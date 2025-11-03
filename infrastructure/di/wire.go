package di

import (
	"app/infrastructure/repository"
	"app/infrastructure/security"
	usecase "app/internal/application/usecase/user"

	"github.com/google/wire"
)

type App struct {
	CreateUserUseCase *usecase.CreateUserUsecase
}

func InitializeApp() *App {
	wire.Build(
		security.NewBcryptPasswordHasher,
		repository.NewUserRepository,
		usecase.NewCreateUserUsecase,
		wire.Struct(new(App), "*"),
	)
	return nil
}
