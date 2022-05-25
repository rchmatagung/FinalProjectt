package entity

import "time"

type User struct {
	ID        int64      `json:"id"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (e *User) TableName() string {
	return "users"
}