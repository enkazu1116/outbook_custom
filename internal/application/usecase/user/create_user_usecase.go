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

// CreateUserUsecase は「ユーザーを新規登録する」というアプリケーションユースケースを表します。
//
// このユースケースの責務は次の通りです。
//   - プレゼンテーション層（ハンドラなど）から受け取った DTO をもとに、
//     ドメインのバリデーションロジックを呼び出す
//   - すでに同じメールアドレスのユーザーが存在しないかリポジトリで確認する
//   - パスワードをドメイン外の PasswordHasher に委譲してハッシュ化する
//   - ドメインエンティティを生成し、リポジトリを通して永続化する
//
// 逆に、「HTTP の詳細」「DB のテーブル構造」「ハッシュアルゴリズムの実装」などには関与しません。
type CreateUserUsecase struct {
	userRepository repository.UserRepository
	hasher         port.PasswordHasher
}

// NewCreateUserUsecase は CreateUserUsecase のコンストラクタです。
// リポジトリと PasswordHasher はポート（インターフェース）越しに注入されるため、
// インフラ層の具体的な実装に依存しないままユースケースをテストできます。
func NewCreateUserUsecase(userRepository repository.UserRepository, hasher port.PasswordHasher) *CreateUserUsecase {
	return &CreateUserUsecase{userRepository: userRepository, hasher: hasher}
}

// CreateUser はユーザー作成ユースケースのエントリポイントです。
// 呼び出し元（ハンドラなど）は DTO とコンテキストを渡すだけで、
// ユースケース内部で以下の一連のフローが実行されます。
//
//  1. ドメインサービスによる入力値のバリデーション
//  2. メールアドレスの重複チェック（UserRepository.ExistsByEmail）
//  3. パスワードのハッシュ化（PasswordHasher.Hash）
//  4. ドメインエンティティの生成（entity.NewUser）
//  5. ユーザーの永続化（UserRepository.CreateUser）
//
// いずれかのステップでエラーが起きた場合は、原因を失わないよう fmt.Errorf(%w) でラップし、
// 呼び出し側で「どこで失敗したか」を追跡しやすいようにしています。
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
