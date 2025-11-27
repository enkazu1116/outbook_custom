package logger

import (
	"testing"
)

// TestLogger はテスト専用のロガーです。
// *testing.T を介してログを出力することで、テスト結果と一緒に確認できるようにします。
type TestLogger struct {
	t *testing.T
}

// New は TestLogger のコンストラクタです。
// 各テストケース内で t を渡して生成して利用します。
func New(t *testing.T) *TestLogger {
	return &TestLogger{t: t}
}

// Info はテスト内での情報ログを出力します。
// メッセージは Value Object で定義したものを利用し、その Message() を渡す想定です。
func (l *TestLogger) Info(msg string, args ...interface{}) {
	l.t.Helper()
	l.t.Logf("[TEST][INFO] "+msg, args...)
}

// Error はテスト内でのエラーログを出力します。
// テスト自体の失敗制御は testing.T に任せ、ここではあくまでログのみを出力します。
func (l *TestLogger) Error(msg string, args ...interface{}) {
	l.t.Helper()
	l.t.Logf("[TEST][ERROR] "+msg, args...)
}


