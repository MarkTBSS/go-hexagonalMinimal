package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/MarkTBSS/go-hexagonalMinimal/adapters"
	"github.com/MarkTBSS/go-hexagonalMinimal/core/employee"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

// Database settings
const (
	host     = "localhost"
	port     = 5432 // Default port
	user     = "postgres"
	password = "pass123"
	dbname   = "hexagonal_minimal"
)

func main() {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Connect with database
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	// Initialize repository and usecase
	repository := adapters.NewDBRepository(db)
	usecase := employee.NewEmployeeUsecase(repository)
	// Initialize HTTP handler
	httpHandler := adapters.NewHTTPHandler(usecase)
	// Create a Fiber app
	app := fiber.New()
	// Routes
	app.Post("/employee", httpHandler.CreateEmployee)
	app.Get("/employee/:id", httpHandler.GetEmployeeByID)
	app.Get("/employees", httpHandler.GetAllEmployees)
	// Listen on port 3000
	log.Fatal(app.Listen(":3000"))
}
