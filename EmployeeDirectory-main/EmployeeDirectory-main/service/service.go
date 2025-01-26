package service

import (
	"employeeeDirectory/models"
	"net/http"
)

type EmployeeService interface {
	CreateEmployee(w http.ResponseWriter, r *http.Request) //Create
	GetEmployee(id int) (models.Employee, error)           //Read
	UpdateEmployee(e models.Employee) error                //Update
	DeleteEmployee(id int) error                           //Delete
	ListAllEmployees()
}
