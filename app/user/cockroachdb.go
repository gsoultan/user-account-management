package user

import "github.com/jinzhu/gorm"

type cockRoachDB struct {
	db interface{}
}

func (c *cockRoachDB) SetPassword(id int64, password []byte) error {
	user, err := c.ReadByID(id)
	if err != nil {
		return err
	}
	return c.db.(*gorm.DB).Model(&user).Update("password", password).Error
}

func (c *cockRoachDB) Create(u *User) error {
	err := c.db.(*gorm.DB).Create(&u).Error
	return err
}

func (c *cockRoachDB) Update(u *User) error {
	user, err := c.ReadByID(*u.ID)
	if err != nil {
		return err
	}

	user.Email = u.Email
	user.UserName = u.UserName
	user.Mobile = u.Mobile

	return c.db.(*gorm.DB).Save(&user).Error
}

func (c *cockRoachDB) DeleteByID(id int64) error {
	user, err := c.ReadByID(id)
	if err != nil {
		return err
	}
	return c.db.(*gorm.DB).Delete(&user).Error
}

func (c *cockRoachDB) ReadAll() ([]User, error) {
	users := []User{}
	err := c.db.(*gorm.DB).Find(&users).Error
	return users, err
}

func (c *cockRoachDB) ReadByID(id int64) (User, error) {
	user := User{}
	err := c.db.(*gorm.DB).First(&user, id).Error
	return user, err
}
