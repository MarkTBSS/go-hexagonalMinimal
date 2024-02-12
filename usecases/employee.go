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
	fmt.Println("usecases/account.go Line 19 : ", repo)
	return &accountUsecase{
		repo: repo,
	}
}

func (u *accountUsecase) CreateAccount(name string) error {
	fmt.Println("usecases/account.go Line 26 : ", name)
	// Implementation details to create account
	return nil
}

func (u *accountUsecase) GetAccountByID(id string) (*repositories.Account, error) {
	fmt.Println("usecases/account.go Line 32 : ", id)
	return u.repo.GetAccountByID(id)
}
