package entity

import (
	"errors"
	"time"
)

// Output Entity
type Output struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	Type        string    `json:"type"`
	Status      string    `json:"status"`
	DeleteFlag  bool      `json:"delete_flag"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// NewOutput コンストラクタ
func NewOutput(userID, title, description, url, outputType string) (*Output, error) {
	// 必須入力チェック（不変的チェック）
	if userID == "" {
		return nil, errors.New("user_id is required")
	}
	if title == "" {
		return nil, errors.New("title is required")
	}

	// Entity生成
	return &Output{
		UserID:      userID,
		Title:       title,
		Description: description,
		URL:         url,
		Type:        outputType,
		Status:      "draft", // デフォルトは下書き
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}
