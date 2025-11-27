package user

import (
	userdto "app/internal/application/dto/user"
	"app/internal/domain/user/entity"
	"app/internal/domain/user/repository"
	"app/internal/domain/user/services"
	"context"
	"fmt"
)

// FindUserUsecase は「ユーザーを条件で検索する」というアプリケーションユースケースを表します。
//
// このユースケースはあくまで「どの条件で検索するか」「検索して得たエンティティをどう返すか」
// というフローを担当し、実際のデータ取得処理は UserRepository に委譲します。
// また、検索条件が妥当かどうかの判断はドメインサービス（FindUserValidation）に集約されています。
type FindUserUsecase struct {
	userRepository repository.UserRepository
}

// NewFindUserUsecase は FindUserUsecase のコンストラクタです。
// ユースケースはリポジトリのインターフェースにのみ依存し、具体的な実装
// （DB, API クライアントなど）はインフラ層でのみ意識されるようにします。
func NewFindUserUsecase(userRepository repository.UserRepository) *FindUserUsecase {
	return &FindUserUsecase{userRepository: userRepository}
}

// FindUser は指定した条件でユーザーを検索するユースケースのメイン処理です。
//
// 処理の流れは次の通りです。
//  1. ドメインサービス FindUserValidation で検索条件（ID/Name/Email）が 1 つ以上指定されているか確認
//  2. 問題がなければリポジトリの FindByUser を呼び出してユーザーを取得
//  3. 取得時に発生したエラーはラップして呼び出し元に返却
//
// これにより、検索条件のルール変更があった場合でもユースケース内の呼び出しは変えずに、
// ドメイン側のバリデーションロジックを変更するだけで済むようになっています。
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
