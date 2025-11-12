//go:build wireinject
// +build wireinject

package di

import (
	"app/infrastructure/db"
	"app/infrastructure/repository"
	"app/infrastructure/security"
	"app/internal/application/port"
	usecase "app/internal/application/usecase/user"

	"github.com/google/wire"
)

type App struct {
	CreateUserUseCase *usecase.CreateUserUsecase
}

func InitializeApp() *App {
	wire.Build(
		db.NewConnection,
		security.NewBcryptPasswordHasher,
		wire.Bind(new(port.PasswordHasher), new(*security.BcryptPasswordHasher)),
		repository.NewUserRepository,
		usecase.NewCreateUserUsecase,
		wire.Struct(new(App), "*"),
	)
	return nil
}
