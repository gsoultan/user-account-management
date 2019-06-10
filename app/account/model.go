package account

import "github.com/gsoultan/user-account-management/app/base"

type Account struct {
	base.Model
	FirstName string `gorm:"index;type:varchar(30);not null"`
	LastName  string `gorm:"index;type:varchar(30)"`
	Company   string `gorm:"index;type:varchar(100);not null"`
	Email     string `gorm:"index;type:varchar(30);not null"`
	Mobile    int    `gorm:"index;not null"`
}
