package user

import "github.com/gsoultan/user-account-management/app/base"

type User struct {
	base.Model
	UserName string `gorm:"type:varchar(30);not null;unique;index"`
	Email    string `gorm:"type:varchar(30);not null;unique;index"`
	Mobile   int    `gorm:"index;not null;unique"`
	Password []byte `gorm:"not null"`
}
