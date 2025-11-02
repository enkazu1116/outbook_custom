package entity

import "time"

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
