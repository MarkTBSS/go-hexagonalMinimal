package main

import (
	"fmt"

	"github.com/MarkTBSS/go-hexagonalMinimal/repositories"
)

func main() {
	repo := repositories.NewAccountRepositories("Charlie")
	fmt.Println("pkg.main.main --> ", repo)
	account, _ := repo.GetAccountByID("z")
	fmt.Println("pkg.main.main --> ", account)
}
