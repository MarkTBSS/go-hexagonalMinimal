package main

import (
	"fmt"

	"github.com/MarkTBSS/go-hexagonalMinimal/repositories"
	"github.com/MarkTBSS/go-hexagonalMinimal/usecases"
)

func main() {
	repo := repositories.NewAccountRepository("Charlie")
	fmt.Println("Repository created: ", repo)

	usecase := usecases.NewAccountUsecase(repo)
	fmt.Println("Usecase created: ", usecase)

	account, _ := usecase.GetAccountByID("z")
	fmt.Println("Account retrieved: ", account)
}
