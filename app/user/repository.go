package user

type Repository interface {
	Create(u *User) error
	Update(u *User) error
	DeleteByID(id int64) error
	SetPassword(id int64, password []byte) error
	ReadAll() ([]User, error)
	ReadByID(id int64) (User, error)
}
