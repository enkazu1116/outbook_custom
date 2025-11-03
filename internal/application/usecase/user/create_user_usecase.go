package user

import (
	user "app/internal/application/dto/user"
	"app/internal/application/port"
	"app/internal/domain/user/entity"
	"app/internal/domain/user/repository"
	"app/internal/domain/user/services"
	"context"
	"fmt"
)

// ユーザー作成ユースケース
type CreateUserUsecase struct {
	userRepository repository.UserRepository
	hasher         port.PasswordHasher
}

// ユーザー作成コンストラクタ
func NewCreateUserUsecase(userRepository repository.UserRepository, hasher port.PasswordHasher) *CreateUserUsecase {
	return &CreateUserUsecase{userRepository: userRepository, hasher: hasher}
}

// ユーザー作成
// 引数: コンテキスト, ユーザー作成コマンド
// 返り値: エラー
// レシーバー: ユーザー作成ユースケースオブジェクト
func (uc *CreateUserUsecase) CreateUser(ctx context.Context, cmd user.CreateUserCommand) error {

	// バリデーションチェック
	if err := services.CreateUserValidation(ctx, cmd.Name, cmd.Email, cmd.Password, cmd.Bio); err != nil {
		return err
	}

	// 重複チェック
	exists, err := uc.userRepository.ExistsByEmail(ctx, cmd.Email)
	// 重複チェックに失敗した場合
	if err != nil {
		return fmt.Errorf("failed to check email duplication: %w", err)
	}
	if exists {
		return fmt.Errorf("email already exists")
	}

	// パスワードのハッシュ化
	hashedPassword, err := uc.hasher.Hash(cmd.Password)

	// パスワードのハッシュ化に失敗した場合
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// Entity生成
	u, err := entity.NewUser(cmd.Name, cmd.Email, hashedPassword, cmd.Bio)
	// Entity生成に失敗した場合
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	// ユーザー作成
	if err := uc.userRepository.CreateUser(ctx, u); err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil

}
