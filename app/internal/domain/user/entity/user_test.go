package entity

import (
	"testing"

	"app/internal/domain/user/value_obj"
	testlogger "app/internal/test/logger"
)

// TestNewUser_Success は NewUser の正常系を検証するテストです。
//
// 将来ユーザーエンティティのフィールドが増えたとしても、
// 「コンストラクタに渡した値が正しく構造体にセットされていること」
// 「作成日時・更新日時がゼロ値でないこと」
// という基本的な性質が守られているかを、ここで素早く確認できるようにしています。
func TestNewUser_Success(t *testing.T) {
	t.Parallel()

	logger := testlogger.New(t)
	logger.Info(value_obj.UserDomainTestStartInfo.Message())
	defer logger.Info(value_obj.UserDomainTestSuccessInfo.Message())

	name := "Alice"
	email := "alice@example.com"
	hashedPassword := "hashed-password"
	bio := "hello"

	u, err := NewUser(name, email, hashedPassword, bio)
	if err != nil {
		t.Fatalf("NewUser() unexpected error: %v", err)
	}

	if u.Name != name {
		t.Errorf("Name = %q, want %q", u.Name, name)
	}
	if u.Email != email {
		t.Errorf("Email = %q, want %q", u.Email, email)
	}
	if u.Password != hashedPassword {
		t.Errorf("Password = %q, want %q", u.Password, hashedPassword)
	}
	if u.Bio != bio {
		t.Errorf("Bio = %q, want %q", u.Bio, bio)
	}
	if u.CreatedAt.IsZero() {
		t.Errorf("CreatedAt is zero, want non-zero time")
	}
	if u.UpdatedAt.IsZero() {
		t.Errorf("UpdatedAt is zero, want non-zero time")
	}
}

// TestNewUser_RequiredFields は必須項目が欠けている場合にエラーとなることを検証します。
//
// name / email / password のいずれかが空文字の場合に、想定しているエラーメッセージが返るかをテーブル形式で確認します。
// これにより、NewUser の入力チェック仕様を変更した場合でも、
// どのパターンでどんなエラーになるべきかをテストから素早く思い出せるようにしています。
func TestNewUser_RequiredFields(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		name          string
		email         string
		hashedPass    string
		wantErrSubstr string
	}{
		"empty name": {
			email:         "alice@example.com",
			hashedPass:    "hashed-password",
			wantErrSubstr: "name is required",
		},
		"empty email": {
			name:          "Alice",
			hashedPass:    "hashed-password",
			wantErrSubstr: "email is required",
		},
		"empty password": {
			name:          "Alice",
			email:         "alice@example.com",
			wantErrSubstr: "password is required",
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			logger := testlogger.New(t)
			logger.Info("NewUser の必須項目テストケース開始: %s", name)

			_, err := NewUser(tt.name, tt.email, tt.hashedPass, "")
			if err == nil {
				t.Fatalf("expected error, got nil")
			}
			if got := err.Error(); got != tt.wantErrSubstr {
				t.Fatalf("error = %q, want %q", got, tt.wantErrSubstr)
			}
		})
	}
}


