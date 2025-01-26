package models

import "fmt"

type Employee struct {
	EmployeeID   int    `json:"id"`
	EmployeeName string `json:"name"`
	EmployeeAge  int    `json:"age"`
	IsMarried    bool   `json:"isMarried"`
}

func (e *Employee) Age() (int, error) {

	return e.EmployeeAge, nil
}

func (e Employee) Ismarried() (bool, error) {

	return e.IsMarried, nil
}

func (e Employee) ID() int {
	return e.EmployeeID
}

func (e Employee) String() string {
	return fmt.Sprintf("Employee ID %v Employee Name %v", e.EmployeeID, e.EmployeeName)
}
