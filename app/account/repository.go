package account

type Repository interface {
	Create(a *Account) error
	Update(a *Account) error
	DeleteByID(id int64) error
	ReadAll() ([]Account, error)
	ReadByID(id int64) (Account, error)
}
