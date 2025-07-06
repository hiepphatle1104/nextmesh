package domain

import "time"

type Session struct {
	ID           string    `gorm:"column:id"`
	AccountID    string    `gorm:"column:account_id"`
	RefreshToken string    `gorm:"column:refresh_token"`
	AccessToken  string    `gorm:"column:access_token"`
	RefreshExp   string    `gorm:"column:refresh_exp"` // 7 days
	ExpiresAt    time.Time `gorm:"column:expires_at"`
	CreatedAt    time.Time `gorm:"column:created_at"`
}

func (Session) TableName() string { return "sessions" }
