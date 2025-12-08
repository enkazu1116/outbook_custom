package logger

import (
	"fmt"
	"log"
	"os"
)

// 日本語ログ出力用の共通ロガー
// 標準出力に日時・ファイル名付きで出力します。
var std = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

// InfoJp は情報レベルの日本語ログを出力します。
func InfoJp(format string, v ...interface{}) {
	std.Printf("[INFO] %s\n", fmt.Sprintf(format, v...))
}

// WarnJp は警告レベルの日本語ログを出力します。
func WarnJp(format string, v ...interface{}) {
	std.Printf("[WARN] %s\n", fmt.Sprintf(format, v...))
}

// ErrorJp はエラーレベルの日本語ログを出力します。
func ErrorJp(format string, v ...interface{}) {
	std.Printf("[ERROR] %s\n", fmt.Sprintf(format, v...))
}

// FatalJp は致命的エラーを日本語で出力し、アプリケーションを終了します。
func FatalJp(format string, v ...interface{}) {
	std.Fatalf("[FATAL] %s\n", fmt.Sprintf(format, v...))
}


