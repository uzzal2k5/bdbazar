package models

import "time"

const (
    RoleSuperAdmin = "superadmin"
    RoleModerator  = "moderator"
    StatusActive    = "active"
    StatusSuspended = "suspended"
)

type Admin struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name" binding:"required"`
	Username  string    `gorm:"uniqueIndex;not null" json:"username" binding:"required"`
	Password  string    `json:"password" binding:"required,min=6"`
	Email     string    `gorm:"uniqueIndex;not null" json:"email" binding:"required,email"`
	Mobile    string    `gorm:"uniqueIndex;not null" json:"mobile" binding:"required,regex=^01[3-9]\\d{8}$"`
	Status    string    `json:"status" gorm:"default:active"` // active, suspended
	Role      string    `gorm:"not null" json:"role" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}