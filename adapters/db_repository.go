package adapters

import (
	"database/sql"
	"errors"

	"github.com/MarkTBSS/go-hexagonalMinimal/core"
)

type DBRepository struct {
	db *sql.DB
}

func NewDBRepository(db *sql.DB) *DBRepository {
	return &DBRepository{db}
}

func (r *DBRepository) Create(employee *core.Employee) error {
	_, err := r.db.Exec("INSERT INTO employees (name, salary, age)VALUES ($1, $2, $3)", employee.Name, employee.Salary, employee.Age)
	return err
}
func (r *DBRepository) GetByID(id string) (*core.Employee, error) {
	var employee core.Employee
	err := r.db.QueryRow("SELECT id, name, salary, age FROM employees WHERE id = $1", id).Scan(&employee.ID, &employee.Name, &employee.Salary, &employee.Age)
	if err != nil {
		return nil, errors.New("employee not found")
	}
	return &employee, nil
}
