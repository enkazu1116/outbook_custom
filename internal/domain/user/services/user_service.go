package services

import (
	"context"
	"regexp"

	"app/internal/domain/user/value_obj"
)

// パスワードの正規表現
var PasswordRegex = regexp.MustCompile(`^[a-zA-Z0-9]+$`)

// CreateUserValidation はユーザー作成時の入力チェックを行います。
func CreateUserValidation(ctx context.Context, name string, email string, password string, bio string) error {

	// 必須入力項目のチェック
	if name == "" || email == "" || password == "" {
		return value_obj.UserRequiredError
	}

	// パスワードの入力数チェック
	// 8文字以上であること
	// 半角英数字であること
	if len(password) < 8 {
		return value_obj.UserPasswordLengthError
	}
	if !PasswordRegex.MatchString(password) {
		return value_obj.UserPasswordFormatError
	}

	// 自己紹介文の入力数チェック
	// 255文字以内であること
	if len(bio) > 255 {
		return value_obj.UserBioLengthError
	}

	// エラーがない場合はnilを返す
	return nil
}

// FindUserValidation はユーザー検索時の入力チェックを行います。
func FindUserValidation(ctx context.Context, id string, name string, email string) error {

	// 検索条件が1つも指定されていない場合はエラー
	if id == "" && name == "" && email == "" {
		return value_obj.UserSearchRequiredError
	}

	return nil
}
