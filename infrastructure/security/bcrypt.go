package security

import "golang.org/x/crypto/bcrypt"

type BcryptPasswordHasher struct{}

// パスワードハッシュ化コンストラクタ
func NewBcryptPasswordHasher() *BcryptPasswordHasher {
	return &BcryptPasswordHasher{}
}

// パスワードハッシュ化
// 引数: 平文パスワード
// 返り値: ハッシュ化されたパスワード, エラー
// レシーバー: パスワードハッシュ化オブジェクト
func (h *BcryptPasswordHasher) Hash(plainPassword string) (string, error) {

	// bcrypt.GenerateFromPassword関数を使用してパスワードをハッシュ化
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)

	return string(bytes), err
}

// パスワード検証
// 引数: 平文パスワード, ハッシュ化されたパスワード
// 返り値: 検証結果
// レシーバー: パスワードハッシュ化オブジェクト
func (h *BcryptPasswordHasher) Compare(plainPassword, hash string) bool {

	// bcrypt.CompareHashAndPassword関数を使用してパスワードを検証
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plainPassword))

	return err == nil
}
