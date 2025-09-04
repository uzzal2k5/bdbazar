package models

import (
    "time"
    "gorm.io/gorm"
)

type Shop struct {
    ID          uint           `gorm:"primaryKey" json:"id"`
    Name        string         `gorm:"unique;not null" json:"name"`
    Description string         `json:"description"`
    OwnerID     uint           `gorm:"index" json:"owner_id"`
    IsApproved  bool           `gorm:"default:false" json:"is_approved"`
    IsBlocked   bool           `gorm:"default:false" json:"is_blocked"`
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
