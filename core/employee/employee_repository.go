package employee

type EmployeeRepository interface {
	Create(employee *Employee) error
	GetByID(id string) (*Employee, error)
	GetAll() ([]*Employee, error)
}
