package account

import (
	"github.com/jinzhu/gorm"
)

type repository struct {
	db interface{}
}

func NewCockroachDBRepository(db interface{}) Repository {
	return &repository{db: db}
}

func (c *repository) Create(a *Account) error {
	err := c.db.(*gorm.DB).Create(&a).Error
	return err
}

func (c *repository) Update(a *Account) error {
	account, err := c.ReadByID(*a.ID)
	if err != nil {
		return err
	}
	account.Email = a.Email
	account.FirstName = a.FirstName
	account.LastName = a.LastName
	account.Mobile = a.Mobile
	account.Company = a.Company
	return c.db.(*gorm.DB).Save(account).Error
}

func (c *repository) DeleteByID(id int64) error {
	account, err := c.ReadByID(id)
	if err != nil {
		return err
	}

	err = c.db.(*gorm.DB).Delete(&account).Error
	return err
}

func (c *repository) ReadAll() ([]Account, error) {
	accounts := []Account{}
	err := c.db.(*gorm.DB).Find(&accounts).Error
	return accounts, err
}

func (c *repository) ReadByID(id int64) (Account, error) {
	account := Account{}
	err := c.db.(*gorm.DB).First(&account, id).Error
	return account, err
}
