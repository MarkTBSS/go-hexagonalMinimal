package main

import (
	"fmt"

	"github.com/MarkTBSS/go-hexagonalMinimal/repositories"
	"github.com/MarkTBSS/go-hexagonalMinimal/usecases"
)

func main() {
	repo := repositories.NewAccountRepository("Charlie")
	fmt.Println("main/main.go Line 12 : ", repo)

	usecase := usecases.NewAccountUsecase(repo)
	fmt.Println("main/main.go Line 15", usecase)

	err := usecase.CreateAccount("Account Name")
	fmt.Println("main/main.go Line 18", err)

	getAccount, _ := usecase.GetAccountByID("2")
	fmt.Println("main/main.go Line 18", getAccount)
}
