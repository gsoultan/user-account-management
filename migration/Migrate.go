package migration

import (
	"github.com/gsoultan/user-account-management/app/account"
	"github.com/gsoultan/user-account-management/app/user"
	"github.com/gsoultan/user-account-management/database"
	"github.com/jinzhu/gorm"
)

func Migrate(d database.Database) {
	d.GetConnection().(*gorm.DB).AutoMigrate(&account.Account{})
	d.GetConnection().(*gorm.DB).AutoMigrate(&user.User{})
}
