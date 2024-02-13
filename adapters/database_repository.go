package adapters

import (
	"database/sql"
	"errors"
	"log"

	"github.com/MarkTBSS/go-hexagonalMinimal/core/employee"
)

type DBRepository struct {
	db *sql.DB
}

func NewDBRepository(db *sql.DB) *DBRepository {
	return &DBRepository{db}
}

func (r *DBRepository) Create(employee *employee.Employee) error {
	log.Println("IN : database_repository.go : adapters.Create()")
	_, err := r.db.Exec("INSERT INTO employees (name, salary, age)VALUES ($1, $2, $3)", employee.Name, employee.Salary, employee.Age)
	log.Println("OUT : database_repository.go : adapters.Create()")
	return err
}
func (r *DBRepository) GetByID(id string) (*employee.Employee, error) {
	var employee employee.Employee
	err := r.db.QueryRow("SELECT id, name, salary, age FROM employees WHERE id = $1", id).Scan(&employee.ID, &employee.Name, &employee.Salary, &employee.Age)
	if err != nil {
		return nil, errors.New("employee not found")
	}
	return &employee, nil
}
func (r *DBRepository) GetAll() ([]*employee.Employee, error) {
	rows, err := r.db.Query("SELECT id, name, salary, age FROM employees")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var employees []*employee.Employee
	for rows.Next() {
		var emp employee.Employee
		err := rows.Scan(&emp.ID, &emp.Name, &emp.Salary, &emp.Age)
		if err != nil {
			return nil, err
		}
		employees = append(employees, &emp)
	}

	return employees, nil
}
