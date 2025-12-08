package port

// パスワードのハッシュ化・検証を行うインターフェース
type PasswordHasher interface {

	// パスワードのハッシュ化
	Hash(password string) (string, error)

	// パスワードの検証
	Compare(password, hash string) bool
}
