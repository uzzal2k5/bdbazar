package models

import "time"

type ActivityLog struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Action    string    `json:"action"`
	ActorID   uint      `json:"actor_id"`
	TargetID  uint      `json:"target_id"`
	Detail    string    `json:"detail"`
	CreatedAt time.Time `json:"created_at"`
}
