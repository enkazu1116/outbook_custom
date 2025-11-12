package services

import (
	"context"
	"errors"
	"regexp"
)

// パスワードの正規表現
var PasswordRegex = regexp.MustCompile(`^[a-zA-Z0-9]+$`)

// エラーメッセージ
var ErrorMessage = map[string]string{
	"required":        "必須入力項目を入力してください。",
	"password_length": "パスワードは8文字以上で入力してください。",
	"password_format": "パスワードは半角英数字で入力してください。",
	"bio_length":      "自己紹介文は255文字以内で入力してください。",
	"search_required": "検索条件を1つ以上指定してください。",
}

// CreateUserValidation はユーザー作成時の入力チェックを行います。
func CreateUserValidation(ctx context.Context, name string, email string, password string, bio string) error {

	// 必須入力項目のチェック
	if name == "" || email == "" || password == "" {
		return errors.New("必須入力項目を入力してください。")
	}

	// パスワードの入力数チェック
	// 8文字以上であること
	// 半角英数字であること
	if len(password) < 8 {
		return errors.New(ErrorMessage["password_length"])
	}
	if !PasswordRegex.MatchString(password) {
		return errors.New(ErrorMessage["password_format"])
	}

	// 自己紹介文の入力数チェック
	// 255文字以内であること
	if len(bio) > 255 {
		return errors.New(ErrorMessage["bio_length"])
	}

	// エラーがない場合はnilを返す
	return nil
}

// FindUserValidation はユーザー検索時の入力チェックを行います。
func FindUserValidation(ctx context.Context, id string, name string, email string) error {

	// 検索条件が1つも指定されていない場合はエラー
	if id == "" && name == "" && email == "" {
		return errors.New(ErrorMessage["search_required"])
	}

	return nil
}
