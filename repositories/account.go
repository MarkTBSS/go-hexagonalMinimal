// repositories/account.go

package repositories

import "fmt"

type AccountRepository interface {
	CreateAccount(a *Account) error
	GetAccountByID(id string) (*Account, error)
}

type Account struct {
	ID   string
	Name string
}

type accountRepository struct {
	dbconn string
}

func NewAccountRepository(dbconn string) AccountRepository {
	fmt.Println("Creating new account repository with connection:", dbconn)
	return &accountRepository{
		dbconn: dbconn,
	}
}

func (r *accountRepository) CreateAccount(a *Account) error {
	fmt.Println("Creating account in repository:", a)
	// Implementation details to create account in database
	return nil
}

func (r *accountRepository) GetAccountByID(id string) (*Account, error) {
	fmt.Println("Fetching account from repository with ID:", id)
	// Implementation details to retrieve account from database
	return &Account{
		ID:   id,
		Name: "Example Account",
	}, nil
}
