package models

import "time"

type SuperAdmin struct {
	ID        string    `gorm:"primaryKey;default:'superadmin'" json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Mobile    string    `json:"mobile"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
