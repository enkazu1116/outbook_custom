package repository

import (
	"app/internal/domain/user/entity"
	"context"
)

// User Entityを扱うRepository
type UserRepository interface {

	// ユーザー作成(全権限使用可能)
	CreateUser(cxt context.Context, user *entity.User) error

	// 登録メールアドレス重複チェック
	ExistsByEmail(cxt context.Context, email string) (bool, error)

	// ユーザー検索(root権限のみ使用可能)
	FindByUser(cxt context.Context, id string, name string, email string) (*entity.User, error)

	// ユーザー更新(root権限のみ使用可能)
	UpdateUser(cxt context.Context, user *entity.User) error

	// ユーザー削除(root権限のみ使用可能)
	DeleteUser(cxt context.Context, id string) error
}
