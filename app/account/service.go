package account

import (
	"encoding/json"
	"github.com/gsoultan/user-account-management/helpers/cache"
	"github.com/jinzhu/gorm"
)

type CommandService interface {
	SignUp(a *Account) error
	Update(a *Account) error
	Remove(id int64) error
}

type QueryService interface {
	ViewAll() ([]Account, error)
	ViewByID(id int64) (Account, error)
}

type TransactionService interface {
	Tx() interface{}
}

type service struct {
	tx                *gorm.DB
	accountRepository Repository
	cache             cache.Redis
}

func (s *service) Tx() interface{} {
	panic("implement me")
}

func NewCommandService(accountRepository Repository, cache cache.Redis) (CommandService, TransactionService) {
	return &service{
			accountRepository: accountRepository,
			cache:             cache,
		}, &service{
			accountRepository: accountRepository,
			cache:             cache,
		}
}

func (s *service) ViewAll() ([]Account, error) {
	accounts := []Account{}
	r, err := s.cache.Exists("accounts")
	if err != nil {
		return accounts, err
	}

	if r == 0 {
		accounts, err = s.accountRepository.ReadAll()
		if err != nil {
			return accounts, err
		}

	}
	rs, err := s.cache.LRange("accounts", 0, -1)
	if err != nil {
		return accounts, err
	}
	err = json.Unmarshal([]byte(rs), &accounts)
	return accounts, err
}

func (s *service) ViewByID(id int64) (Account, error) {
	return s.accountRepository.ReadByID(id)
}

func (s *service) SignUp(a *Account) error {
	err := s.accountRepository.Create(a)
	return err
}

func (s *service) Update(a *Account) error {
	err := s.accountRepository.Update(a)
	return err
}

func (s *service) Remove(id int64) error {
	err := s.accountRepository.DeleteByID(id)
	return err
}
