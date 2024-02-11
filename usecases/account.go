package usecases

import "github.com/MarkTBSS/go-hexagonalMinimal/repositories"

type CreateAccountData struct {
	Name    string
	Balance float32
}

type AccountData struct {
	ID   string
	Name string
}

type AccountUsecase interface {
	CreateAccount(a *CreateAccountData) error
	GetAccountByID(id string) (*AccountData, error)
}

type accountUsecase struct {
	repo repositories.AccountRepositories
}

func NewAccountUsecase(repo repositories.AccountRepositories) AccountUsecase {
	return &accountUsecase{
		repo: repo,
	}
}

func (u *accountUsecase) CreateAccount(a *CreateAccountData) error {
	return nil
}

func (u *accountUsecase) GetAccountByID(id string) (*AccountData, error) {
	return nil, nil
}
