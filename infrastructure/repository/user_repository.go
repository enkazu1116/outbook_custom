package repository

import (
	userEntity "app/internal/domain/user/entity"
	userRepository "app/internal/domain/user/repository"
	"context"
	"errors"
	"strings"

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

// CreateUser はユーザーを新規登録します。
// 引数: コンテキスト, 登録するユーザーエンティティ
// 返り値: 永続化に失敗した場合はエラー
// レシーバー: ユーザーリポジトリオブジェクト
func (r *UserRepositoryImpl) CreateUser(cxt context.Context, user *userEntity.User) error {

	return r.db.WithContext(cxt).Create(user).Error
}

// ExistsByEmail はメールアドレスの重複を確認します。
// 引数: コンテキスト, 確認対象のメールアドレス
// 返り値: true=重複あり, false=重複なし, 判定に失敗した場合はエラー
// レシーバー: ユーザーリポジトリオブジェクト
func (r *UserRepositoryImpl) ExistsByEmail(cxt context.Context, email string) (bool, error) {

	var count int64

	// メールアドレスが一致するEntityの数を取得
	if err := r.db.WithContext(cxt).
		Model(&userEntity.User{}).
		Where("email = ? AND delete_flag = ?", email, false).
		Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

// FindByUser は指定した条件のいずれかに一致するユーザーを取得します。
// 引数: コンテキスト, 検索条件としてのID/名前/メールアドレス（空文字は無視）
// 返り値: 一致したユーザー, 見つからない/検索に失敗した場合はエラー
// レシーバー: ユーザーリポジトリオブジェクト
func (r *UserRepositoryImpl) FindByUser(cxt context.Context, id string, name string, email string) (*userEntity.User, error) {

	conditions := make([]string, 0, 4)
	values := make([]interface{}, 0, 4)

	if id != "" {
		conditions = append(conditions, "id = ?")
		values = append(values, id)
	}
	if name != "" {
		conditions = append(conditions, "name = ?")
		values = append(values, name)
	}
	if email != "" {
		conditions = append(conditions, "email = ?")
		values = append(values, email)
	}

	// 論理削除されていないユーザーのみ対象
	conditions = append(conditions, "delete_flag = ?")
	values = append(values, false)

	if len(conditions) == 0 {
		return nil, errors.New("no search criteria provided")
	}

	var u userEntity.User
	err := r.db.WithContext(cxt).
		Model(&userEntity.User{}).
		Where(strings.Join(conditions, " OR "), values...).
		First(&u).Error
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// UpdateUser は既存ユーザー情報を更新します。
// 引数: コンテキスト, 更新後のユーザーエンティティ（ID必須）
// 返り値: 更新に失敗した場合はエラー
// レシーバー: ユーザーリポジトリオブジェクト
func (r *UserRepositoryImpl) UpdateUser(cxt context.Context, user *userEntity.User) error {
	return r.db.WithContext(cxt).Model(&userEntity.User{}).Where("id = ?", user.ID).Updates(user).Error
}

// DeleteUser は指定したユーザーを削除します。
// 引数: コンテキスト, 削除対象ID
// 返り値: 削除に失敗した場合はエラー
// レシーバー: ユーザーリポジトリオブジェクト
func (r *UserRepositoryImpl) DeleteUser(cxt context.Context, id string) error {
	// 物理削除ではなく論理削除（delete_flag を立てる）のみに変更
	return r.db.WithContext(cxt).
		Model(&userEntity.User{}).
		Where("id = ?", id).
		Update("delete_flag", true).Error
}
