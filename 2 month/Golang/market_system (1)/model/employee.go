package model

type Employee struct {
	Id       string
	DateTime string
	Employee string
	Adresss  string
	Status   string
}

type EmployeePrimaryKey struct {
	Id string
}

type EmployeeUpdate struct {
	Id       string
	DateTime string
	Employee string
	Adresss  string
	Status   string
}

type GetListEmployeeRequest struct {
	Offset int
	Limit  int
}

type GetListEmployeeResponse struct {
	Count   int
	Employees []Employee
}
