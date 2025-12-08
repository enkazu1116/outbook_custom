package entity

import (
	"errors"
	"time"
)

// User Entity
type User struct {
	ID                string    `json:"id"`
	Name              string    `json:"name"`
	Email             string    `json:"email"`
	Password          string    `json:"password"`
	Role              string    `json:"role"`
	Bio               string    `json:"bio"`
	SkillLevel        string    `json:"skill_level"`
	YearsOfExperience int       `json:"years_of_experience"`
	DeleteFlag        bool      `json:"delete_flag"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// NewUser コンストラクタ
func NewUser(name, email, hashedPassword, bio string) (*User, error) {
	// 必須入力チェック（不変的チェック）
	if name == "" {
		return nil, errors.New("name is required")
	}
	if email == "" {
		return nil, errors.New("email is required")
	}
	if hashedPassword == "" {
		return nil, errors.New("password is required")
	}

	// Entity生成
	return &User{
		Name:      name,
		Email:     email,
		Password:  hashedPassword,
		Bio:       bio,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
