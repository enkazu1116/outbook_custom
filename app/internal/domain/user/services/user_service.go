package services

import (
	"context"
	"regexp"

	"app/internal/domain/user/value_obj"
)

// PasswordRegex はパスワード文字列の形式チェックに使用する正規表現です。
// 「半角英大文字・半角英小文字・数字」のみを許可し、その他の記号や全角文字を弾くことで
// 誤入力や意図しない文字種を早期に検知するために利用します。
var PasswordRegex = regexp.MustCompile(`^[a-zA-Z0-9]+$`)

// CreateUserValidation は「ユーザーを新規作成してよい状態か」を判定するためのドメインバリデーションです。
// アプリケーション層やハンドラ層から呼び出され、以下のルールを一括でチェックします。
//
//   - 必須入力: name, email, password のいずれかが欠けていればエラー
//   - パスワード長: 8文字未満であればエラー
//   - パスワード形式: 半角英数字以外の文字を含んでいればエラー
//   - 自己紹介文: 255文字を超えていればエラー
//
// これらのルールはユーザードメインに閉じたビジネスルールであり、呼び出し側からは
// value_obj で定義されたエラーメッセージを通して「何が原因で作成できないのか」を把握できます。
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

// FindUserValidation はユーザー検索時に「検索条件がまったく指定されていない」状態を防ぐためのドメインバリデーションです。
// ID / Name / Email のいずれか 1 つでも値が入っていれば検索を許可し、
// すべて空文字の場合は「検索条件を1つ以上指定してください」というメッセージを返します。
// どのレイヤから利用しても同じルールで検索条件をチェックできるように、この関数にロジックを集約しています。
func FindUserValidation(ctx context.Context, id string, name string, email string) error {

	// 検索条件が1つも指定されていない場合はエラー
	if id == "" && name == "" && email == "" {
		return value_obj.UserSearchRequiredError
	}

	return nil
}
