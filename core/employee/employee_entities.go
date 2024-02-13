package employee

type Employee struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Salary string `json:"salary"`
	Age    string `json:"age"`
}

type Employees struct {
	Employees []Employee `json:"employees"`
}
