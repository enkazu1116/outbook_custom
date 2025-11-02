package value_obj

type Role string

// 権限定義
const (
	Root   Role = "root"
	Admin  Role = "admin"
	Member Role = "member"
	Guest  Role = "guest"
)
