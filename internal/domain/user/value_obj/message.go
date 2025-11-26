package value_obj

// MessageLevel はメッセージの重要度を表すレベルです。
type MessageLevel string

const (
	MessageLevelError MessageLevel = "error"
	MessageLevelWarn  MessageLevel = "warn"
	MessageLevelInfo  MessageLevel = "info"
)

// DomainMessage はドメイン内で扱うメッセージの共通インターフェースです。
type DomainMessage interface {
	Level() MessageLevel
	Code() string
	Message() string
}

// ErrorMessage はエラーメッセージを表す値オブジェクトです。
// error インターフェースも実装しているため、そのまま error として返却できます。
type ErrorMessage struct {
	code    string
	message string
}

func (m ErrorMessage) Level() MessageLevel {
	return MessageLevelError
}

func (m ErrorMessage) Code() string {
	return m.code
}

func (m ErrorMessage) Message() string {
	return m.message
}

// Error implements the error interface.
func (m ErrorMessage) Error() string {
	return m.message
}

// WarnMessage は警告レベルのメッセージを表します。
type WarnMessage struct {
	code    string
	message string
}

func (m WarnMessage) Level() MessageLevel {
	return MessageLevelWarn
}

func (m WarnMessage) Code() string {
	return m.code
}

func (m WarnMessage) Message() string {
	return m.message
}

// InfoMessage は情報レベルのメッセージを表します。
type InfoMessage struct {
	code    string
	message string
}

func (m InfoMessage) Level() MessageLevel {
	return MessageLevelInfo
}

func (m InfoMessage) Code() string {
	return m.code
}

func (m InfoMessage) Message() string {
	return m.message
}

// --- User ドメイン向けのメッセージ定義 ---

var (
	// 必須入力系
	UserRequiredError = ErrorMessage{
		code:    "user.required",
		message: "必須入力項目を入力してください。",
	}

	// パスワード関連
	UserPasswordLengthError = ErrorMessage{
		code:    "user.password.length",
		message: "パスワードは8文字以上で入力してください。",
	}
	UserPasswordFormatError = ErrorMessage{
		code:    "user.password.format",
		message: "パスワードは半角英数字で入力してください。",
	}

	// 自己紹介関連
	UserBioLengthError = ErrorMessage{
		code:    "user.bio.length",
		message: "自己紹介文は255文字以内で入力してください。",
	}

	// 検索関連
	UserSearchRequiredError = ErrorMessage{
		code:    "user.search.required",
		message: "検索条件を1つ以上指定してください。",
	}
)


