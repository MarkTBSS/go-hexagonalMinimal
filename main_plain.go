package main

/* package main

import (
    "database/sql"
    "fmt"
    "log"

    "github.com/gofiber/fiber/v2"
    _ "github.com/lib/pq"
)

// Database instance
var db *sql.DB

    "github.com/MarkTBSS/go-hexagonalMinimal/repositories"
    "github.com/MarkTBSS/go-hexagonalMinimal/usecases"
// Database settings
const (
    host     = "localhost"
    port     = 5432 // Default port
    user     = "postgres"
    password = "pass123"
    dbname   = "hexagonal_minimal"
)

// Employee struct
type Employee struct {
    ID     int64  `json:"id"`
    Name   string `json:"name"`
    Salary string `json:"salary"`
    Age    string `json:"age"`
}

// Employees struct
type Employees struct {
    Employees []Employee `json:"employees"`
}

func Connect() error {
    var err error
    db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
    if err != nil {
        return err
    }
    if err = db.Ping(); err != nil {
        return err
    }
    return nil
}

func main() {
    repo := repositories.NewAccountRepository("Charlie")
    fmt.Println("main/main.go Line 12 : ", repo)
    // Connect with database
    if err := Connect(); err != nil {
        log.Fatal(err)
    }

    // Create a Fiber app
    app := fiber.New()

    // Get all records from postgreSQL
    app.Get("/employee", func(c *fiber.Ctx) error {
        // Select all Employee(s) from database
        rows, err := db.Query("SELECT id, name, salary, age FROM employees order by id")
        if err != nil {
            return c.Status(500).SendString(err.Error())
        }
        defer rows.Close()
        result := Employees{}

        for rows.Next() {
            employee := Employee{}
            if err := rows.Scan(&employee.ID, &employee.Name, &employee.Salary, &employee.Age); err != nil {
                return err // Exit if we get an error
            }

            // Append Employee to Employees
            result.Employees = append(result.Employees, employee)
        }
        // Return Employees in JSON format
        return c.JSON(result)
    })

    // Add record into postgreSQL
    app.Post("/employee", func(c *fiber.Ctx) error {
        // New Employee struct
        u := new(Employee)

        // Parse body into struct
        if err := c.BodyParser(u); err != nil {
            return c.Status(400).SendString(err.Error())
        }

        // Insert Employee into database
        res, err := db.Query("INSERT INTO employees (name, salary, age)VALUES ($1, $2, $3)", u.Name, u.Salary, u.Age)
        if err != nil {
            return err
        }

        // Print result
        log.Println(res)

        // Return Employee in JSON format
        return c.JSON(u)
    })

    // Update record into postgreSQL
    app.Put("/employee", func(c *fiber.Ctx) error {
        // New Employee struct
        u := new(Employee)

        // Parse body into struct
        if err := c.BodyParser(u); err != nil {
            return c.Status(400).SendString(err.Error())
        }

        // Update Employee into database
        res, err := db.Query("UPDATE employees SET name=$1,salary=$2,age=$3 WHERE id=$4", u.Name, u.Salary, u.Age, u.ID)
        if err != nil {
            return err
        }

    usecase := usecases.NewAccountUsecase(repo)
    fmt.Println("main/main.go Line 15", usecase)
        // Print result
        log.Println(res)

    err := usecase.CreateAccount("Account Name")
    fmt.Println("main/main.go Line 18", err)
        // Return Employee in JSON format
        return c.Status(201).JSON(u)
    })

    getAccount, _ := usecase.GetAccountByID("2")
    fmt.Println("main/main.go Line 18", getAccount)
    log.Fatal(app.Listen(":3000"))
} */
