package models

import (
    "time"
    "gorm.io/gorm"
    "gorm.io/datatypes"
)


type User struct {
    ID        uint           `gorm:"primaryKey" json:"id"`
    Name      string         `json:"name"`
    Email     string         `gorm:"uniqueIndex" json:"email"`
    Mobile    string         `gorm:"uniqueIndex" json:"mobile"`
    Password  string         `json:"-"`
    Roles     datatypes.JSON `json:"roles"`
    RefreshTokens []RefreshToken `gorm:"foreignKey:UserID"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}



type RefreshToken struct {
    ID        uint      `gorm:"primaryKey"`
    UserID    uint      `gorm:"not null"`
    Token     string    `gorm:"uniqueIndex;not null"`
    ExpiresAt time.Time `gorm:"not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
}
