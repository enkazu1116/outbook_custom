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

// --- Output ドメイン向けのメッセージ定義 ---

var (
	// 必須入力系
	OutputRequiredError = ErrorMessage{
		code:    "output.required",
		message: "必須入力項目を入力してください。",
	}

	// タイトル関連
	OutputTitleLengthError = ErrorMessage{
		code:    "output.title.length",
		message: "タイトルは255文字以内で入力してください。",
	}

	// 説明関連
	OutputDescriptionLengthError = ErrorMessage{
		code:    "output.description.length",
		message: "説明文は1000文字以内で入力してください。",
	}

	// URL関連
	OutputURLFormatError = ErrorMessage{
		code:    "output.url.format",
		message: "URLの形式が正しくありません。",
	}

	// --- テスト用メッセージ ---

	// OutputDomainTestStartInfo はアウトプットドメイン層のテスト開始を表す情報メッセージです。
	OutputDomainTestStartInfo = InfoMessage{
		code:    "test.output.domain.start",
		message: "アウトプットドメイン層のテストを開始します。",
	}

	// OutputDomainTestSuccessInfo はアウトプットドメイン層のテスト成功を表す情報メッセージです。
	OutputDomainTestSuccessInfo = InfoMessage{
		code:    "test.output.domain.success",
		message: "アウトプットドメイン層のテストが正常に完了しました。",
	}

	// OutputUsecaseTestStartInfo はアウトプットユースケース層のテスト開始を表す情報メッセージです。
	OutputUsecaseTestStartInfo = InfoMessage{
		code:    "test.output.usecase.start",
		message: "アウトプットユースケース層のテストを開始します。",
	}

	// OutputUsecaseTestSuccessInfo はアウトプットユースケース層のテスト成功を表す情報メッセージです。
	OutputUsecaseTestSuccessInfo = InfoMessage{
		code:    "test.output.usecase.success",
		message: "アウトプットユースケース層のテストが正常に完了しました。",
	}
)
