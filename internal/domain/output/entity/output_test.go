package entity

import (
	"testing"

	"app/internal/domain/output/value_obj"
	testlogger "app/internal/test/logger"
)

// TestNewOutput_Success は NewOutput の正常系を検証するテストです。
//
// 将来アウトプットエンティティのフィールドが増えたとしても、
// 「コンストラクタに渡した値が正しく構造体にセットされていること」
// 「作成日時・更新日時がゼロ値でないこと」
// という基本的な性質が守られているかを、ここで素早く確認できるようにしています。
func TestNewOutput_Success(t *testing.T) {
	t.Parallel()

	logger := testlogger.New(t)
	logger.Info(value_obj.OutputDomainTestStartInfo.Message())
	defer logger.Info(value_obj.OutputDomainTestSuccessInfo.Message())

	userID := "user-123"
	title := "テストアウトプット"
	description := "これはテスト用のアウトプットです"
	url := "https://example.com/output"
	outputType := "blog"

	o, err := NewOutput(userID, title, description, url, outputType)
	if err != nil {
		t.Fatalf("NewOutput() unexpected error: %v", err)
	}

	if o.UserID != userID {
		t.Errorf("UserID = %q, want %q", o.UserID, userID)
	}
	if o.Title != title {
		t.Errorf("Title = %q, want %q", o.Title, title)
	}
	if o.Description != description {
		t.Errorf("Description = %q, want %q", o.Description, description)
	}
	if o.URL != url {
		t.Errorf("URL = %q, want %q", o.URL, url)
	}
	if o.Type != outputType {
		t.Errorf("Type = %q, want %q", o.Type, outputType)
	}
	if o.Status != "draft" {
		t.Errorf("Status = %q, want %q", o.Status, "draft")
	}
	if o.CreatedAt.IsZero() {
		t.Errorf("CreatedAt is zero, want non-zero time")
	}
	if o.UpdatedAt.IsZero() {
		t.Errorf("UpdatedAt is zero, want non-zero time")
	}
}

// TestNewOutput_RequiredFields は必須項目が欠けている場合にエラーとなることを検証します。
//
// user_id / title のいずれかが空文字の場合に、想定しているエラーメッセージが返るかをテーブル形式で確認します。
// これにより、NewOutput の入力チェック仕様を変更した場合でも、
// どのパターンでどんなエラーになるべきかをテストから素早く思い出せるようにしています。
func TestNewOutput_RequiredFields(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		userID        string
		title         string
		description   string
		url           string
		outputType    string
		wantErrSubstr string
	}{
		"empty user_id": {
			userID:        "",
			title:         "テストアウトプット",
			description:   "これはテスト用のアウトプットです",
			url:           "https://example.com/output",
			outputType:    "blog",
			wantErrSubstr: "user_id is required",
		},
		"empty title": {
			userID:        "user-123",
			title:         "",
			description:   "これはテスト用のアウトプットです",
			url:           "https://example.com/output",
			outputType:    "blog",
			wantErrSubstr: "title is required",
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			logger := testlogger.New(t)
			logger.Info("NewOutput の必須項目テストケース開始: %s", name)

			_, err := NewOutput(tt.userID, tt.title, tt.description, tt.url, tt.outputType)
			if err == nil {
				t.Fatalf("expected error, got nil")
			}
			if got := err.Error(); got != tt.wantErrSubstr {
				t.Fatalf("error = %q, want %q", got, tt.wantErrSubstr)
			}
		})
	}
}
