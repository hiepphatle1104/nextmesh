package model

import "time"

type Session struct {
	ID        string    `gorm:"column:id;"`
	AccountID string    `gorm:"column:account_id;"`
	ExpiresAt time.Time `gorm:"column:expires_at;"`
	CreatedAt time.Time `gorm:"column:created_at;"`
}
