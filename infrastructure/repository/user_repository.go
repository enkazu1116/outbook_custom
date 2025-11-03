package repository

import (
	userEntity "app/internal/domain/user/entity"
	userRepository "app/internal/domain/user/repository"
	"context"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

// ユーザーリポジトリコンストラクタ
// 引数: データベースオブジェクト
// 返り値: ユーザーリポジトリオブジェクト
func NewUserRepository(db *gorm.DB) userRepository.UserRepository {
	return &UserRepositoryImpl{db: db}
}

// リポジトリ実装層
// ユーザー作成
// 引数: コンテキスト, ユーザーエンティティ
// 返り値: エラー
// レシーバー: ユーザーリポジトリオブジェクト
func (r *UserRepositoryImpl) CreateUser(cxt context.Context, user *userEntity.User) error {

	return r.db.WithContext(cxt).Create(user).Error
}

// 登録メールアドレス重複チェック
// 引数: コンテキスト, メールアドレス
// 返り値: 重複チェック結果, エラー
// レシーバー: ユーザーリポジトリオブジェクト
func (r *UserRepositoryImpl) ExistsByEmail(cxt context.Context, email string) (bool, error) {

	var count int64

	// メールアドレスが一致するEntityの数を取得
	if err := r.db.WithContext(cxt).Model(&userEntity.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

// ユーザー検索
// 引数: コンテキスト, id, name, email
// 返り値: ユーザーエンティティ, エラー
// レシーバー: ユーザーリポジトリオブジェクト
func (r *UserRepositoryImpl) FindByUser(cxt context.Context, id string, name string, email string) (*userEntity.User, error) {

	u := &userEntity.User{}

	// ユーザー検索
	// 検索条件: id, name, emailいずれかが一致するユーザーを取得
	err := r.db.WithContext(cxt).
		Model(&userEntity.User{}).
		Where("id = ? OR name = ? OR email = ?", id, name, email).
		First(u).Error

	// エラーが発生した場合はnilを返す
	if err != nil {
		return nil, err
	}
	return u, nil
}

// ユーザー更新
func (r *UserRepositoryImpl) UpdateUser(cxt context.Context, user *userEntity.User) error {
	return r.db.WithContext(cxt).Model(&userEntity.User{}).Where("id = ?", user.ID).Updates(user).Error
}

// ユーザー削除
func (r *UserRepositoryImpl) DeleteUser(cxt context.Context, id string) error {
	return r.db.WithContext(cxt).Model(&userEntity.User{}).Where("id = ?", id).Delete(&userEntity.User{}).Error
}
