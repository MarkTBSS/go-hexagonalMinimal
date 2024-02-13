package employee

import "log"

type EmployeeUsecase struct {
	repository EmployeeRepository
}

func NewEmployeeUsecase(repository EmployeeRepository) *EmployeeUsecase {
	return &EmployeeUsecase{repository}
}

func (uc *EmployeeUsecase) CreateEmployee(name, salary, age string) error {
	log.Println("IN : employee.CreateEmployee()")
	Employee := &Employee{
		Name:   name,
		Salary: salary,
		Age:    age,
	}
	log.Println("OUT : employee.CreateEmployee()")
	return uc.repository.Create(Employee)
}

func (uc *EmployeeUsecase) GetEmployeeByID(id string) (*Employee, error) {
	return uc.repository.GetByID(id)
}

func (uc *EmployeeUsecase) GetAllEmployees() ([]*Employee, error) {
	return uc.repository.GetAll()
}
