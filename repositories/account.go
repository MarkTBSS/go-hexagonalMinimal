package repositories

import "fmt"

func NewAccountRepositories(dbconn string) AccountRepositories {
	fmt.Println("pkg.repositories.account.6 --> ", dbconn)
	return &accountRepositories{
		dbconn: dbconn,
	}
}

// ประกาศให้รู้ว่าเราจะทำอะไรได้บ้าง (ไม่รวม Business Logic)
type AccountRepositories interface {
	CreateAccount(a *Account) error
	GetAccountByID(id string) (*Account, error)
}

// Database Model or Entities
type Account struct {
	ID   string
	Name string
}

// Database Connection
type accountRepositories struct {
	dbconn string
}

// Repository Constructor
func (r *accountRepositories) CreateAccount(a *Account) error {
	fmt.Println("pkg.repositories.account.CreateAccount()")
	return nil
}

func (r *accountRepositories) GetAccountByID(id string) (*Account, error) {
	fmt.Println("pkg.repositories.account.GetAccountByID()")
	if id == "a" {
		return &Account{
			ID:   "a",
			Name: "a",
		}, nil
	} else if id == "z" {
		return &Account{
			ID:   "z",
			Name: "z",
		}, nil
	}
	return nil, nil
}
