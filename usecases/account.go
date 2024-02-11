// usecases/account.go

package usecases

import (
	"fmt"

	"github.com/MarkTBSS/go-hexagonalMinimal/repositories"
)

type AccountUsecase interface {
	CreateAccount(name string) error
	GetAccountByID(id string) (*repositories.Account, error)
}

type accountUsecase struct {
	repo repositories.AccountRepository
}

func NewAccountUsecase(repo repositories.AccountRepository) AccountUsecase {
	fmt.Println("Creating new account usecase with repository:", repo)
	return &accountUsecase{
		repo: repo,
	}
}

func (u *accountUsecase) CreateAccount(name string) error {
	fmt.Println("Creating account with name:", name)
	// Implementation details to create account
	return nil
}

func (u *accountUsecase) GetAccountByID(id string) (*repositories.Account, error) {
	fmt.Println("Fetching account with ID:", id)
	return u.repo.GetAccountByID(id)
}
