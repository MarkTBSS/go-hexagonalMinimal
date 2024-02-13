package core

type EmployeeRepository interface {
	Create(employee *Employee) error
	GetByID(id string) (*Employee, error)
}
