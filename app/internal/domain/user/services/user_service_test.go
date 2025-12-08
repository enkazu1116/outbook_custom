package services

import (
	"context"
	"testing"

	"app/internal/domain/user/value_obj"
	testlogger "app/internal/test/logger"
)

// TestFindUserValidation はユーザー検索時の入力チェックの動作を検証します。
//
// 「どの検索条件の組み合わせならエラーにならないのか / なるのか」を一覧で確認できるようにすることで、
// FindUserValidation の仕様を後から読み返したときにイメージしやすくしています。
func TestFindUserValidation(t *testing.T) {
	t.Parallel()

	logger := testlogger.New(t)
	logger.Info(value_obj.UserDomainTestStartInfo.Message())
	defer logger.Info(value_obj.UserDomainTestSuccessInfo.Message())

	tests := map[string]struct {
		id    string
		name  string
		email string
		err   error
	}{
		"all empty": {
			err: value_obj.UserSearchRequiredError,
		},
		"id only": {
			id: "user-1",
		},
		"name only": {
			name: "Alice",
		},
		"email only": {
			email: "alice@example.com",
		},
		"multiple fields": {
			id:    "user-1",
			name:  "Alice",
			email: "alice@example.com",
		},
	}

	ctx := context.Background()

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			logger := testlogger.New(t)
			logger.Info("FindUserValidation テストケース開始: %s", name)

			err := FindUserValidation(ctx, tt.id, tt.name, tt.email)

			switch {
			case tt.err == nil && err != nil:
				t.Fatalf("unexpected error: %v", err)
			case tt.err != nil && err == nil:
				t.Fatalf("expected error %v, got nil", tt.err)
			case tt.err != nil && err != nil && err.Error() != tt.err.Error():
				t.Fatalf("expected error %v, got %v", tt.err, err)
			}
		})
	}
}

// TestCreateUserValidation はユーザー作成時の入力チェックの動作を検証します。
//
// 必須入力・パスワード長・パスワード形式・自己紹介文の長さなど、
// CreateUserValidation に閉じ込められたドメインルールが正しく機能しているかを
// ケースごとに表形式で確認します。
// どのパターンでどのエラーを返すかを、このテストを読むだけで一望できるようにすることが狙いです。
func TestCreateUserValidation(t *testing.T) {
	t.Parallel()

	logger := testlogger.New(t)
	logger.Info(value_obj.UserDomainTestStartInfo.Message())
	defer logger.Info(value_obj.UserDomainTestSuccessInfo.Message())

	ctx := context.Background()

	tests := map[string]struct {
		name     string
		email    string
		password string
		bio      string
		wantErr  error
	}{
		"all required empty": {
			wantErr: value_obj.UserRequiredError,
		},
		"password too short": {
			name:     "Alice",
			email:    "alice@example.com",
			password: "short",
			wantErr:  value_obj.UserPasswordLengthError,
		},
		"password invalid format": {
			name:     "Alice",
			email:    "alice@example.com",
			password: "invalid!",
			wantErr:  value_obj.UserPasswordFormatError,
		},
		"bio too long": {
			name:     "Alice",
			email:    "alice@example.com",
			password: "Password1",
			bio:      string(make([]byte, 256)),
			wantErr:  value_obj.UserBioLengthError,
		},
		"valid input": {
			name:     "Alice",
			email:    "alice@example.com",
			password: "Password1",
			bio:      "hello",
			wantErr:  nil,
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			logger := testlogger.New(t)
			logger.Info("CreateUserValidation テストケース開始: %s", name)

			err := CreateUserValidation(ctx, tt.name, tt.email, tt.password, tt.bio)

			switch {
			case tt.wantErr == nil && err != nil:
				t.Fatalf("unexpected error: %v", err)
			case tt.wantErr != nil && err == nil:
				t.Fatalf("expected error %v, got nil", tt.wantErr)
			case tt.wantErr != nil && err != nil && err.Error() != tt.wantErr.Error():
				t.Fatalf("expected error %v, got %v", tt.wantErr, err)
			}
		})
	}
}
