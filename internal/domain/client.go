package domain

import "time"

type Client struct {
	ID          string    `gorm:"column:id"`
	SecretKey   string    `gorm:"column:secret_key"`
	RedirectURI string    `gorm:"column:redirect_uri"`
	AccountID   string    `gorm:"column:account_id"`
	CreatedAt   time.Time `gorm:"column:created_at"`
}

func (Client) TableName() string { return "clients" }
