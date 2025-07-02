package model

import "time"

type Account struct {
	ID        string    `gorm:"column:id;"`
	Name      string    `gorm:"column:name;"`
	Email     string    `gorm:"column:email;"`
	Password  string    `gorm:"column:password;"`
	CreatedAt time.Time `gorm:"column:created_at;"`
}

func (Account) TableName() string { return "accounts" }
