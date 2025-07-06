package domain

type Account struct {
	ID       string `gorm:"column:id"`
	Name     string `gorm:"column:name"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}

func (Account) TableName() string { return "accounts" }
