package core

type Employee struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Salary string `json:"salary"`
	Age    string `json:"age"`
}

type Employees struct {
	Employees []Employee `json:"employees"`
}

type EmployeeUsecase struct {
	repository EmployeeRepository
}

func NewEmployeeUsecase(repository EmployeeRepository) *EmployeeUsecase {
	return &EmployeeUsecase{repository}
}

func (uc *EmployeeUsecase) CreateEmployee(name, salary, age string) error {
	Employee := &Employee{
		Name:   name,
		Salary: salary,
		Age:    age,
	}
	return uc.repository.Create(Employee)
}

func (uc *EmployeeUsecase) GetEmployeeByID(id string) (*Employee, error) {
	return uc.repository.GetByID(id)
}
