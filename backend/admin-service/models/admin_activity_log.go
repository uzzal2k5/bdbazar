package models

import "time"

type AdminActivityLog struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    AdminID   uint      `json:"admin_id"`             // refers to Admin
    Action    string    `json:"action"`               // e.g., "approved_shop", "banned_user"
    Target    string    `json:"target"`               // e.g., "shop", "user"
    TargetID  uint      `json:"target_id"`            // ID from shop-service or auth-service
    Note      string    `json:"note,omitempty"`       // optional comment
    CreatedAt time.Time `json:"created_at"`
    LoggedAt  time.Time `json:"logged_at"` // ✅ Add this line
}
