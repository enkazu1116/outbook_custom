package value_obj

type Role string

// 権限定義
const (
	Root   Role = "root"
	Admin  Role = "admin"
	Member Role = "member"
	Guest  Role = "guest"
)

// ゲストユーザーの規制メソッド
func (r Role) IsMember() bool {

	switch r {
	case Guest:
		return false
	case Member:
		return true
	case Admin:
		return true
	case Root:
		return true
	}
	return false
}

// 管理者権限のチェック
func (r Role) IsAdmin() bool {
	switch r {
	case Admin:
		return true
	case Root:
		return true
	}
	return false
}

// ルート権限のチェック
func (r Role) IsRoot() bool {
	switch r {
	case Root:
		return true
	}
	return false
}
