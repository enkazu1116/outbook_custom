package user

// CreateUserCommand はユーザー作成時の入力データを保持します。
type CreateUserCommand struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Bio      string `json:"bio"`
}

// FindUserQuery はユーザー検索時の条件を表します。
type FindUserQuery struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
