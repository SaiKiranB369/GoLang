package repository

import (
	"employeeeDirectory/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type EmployeeRepo struct {
	employees map[int]models.Employee
}

func NewEmployeeRepo() *EmployeeRepo {
	return &EmployeeRepo{employees: make(map[int]models.Employee)}
}

func (r *EmployeeRepo) CreateEmployee(w http.ResponseWriter, req *http.Request) {

	var emp models.Employee

	if err := json.NewDecoder(req.Body).Decode(&emp); err != nil {
		http.Error(w, "Invalid Request body", http.StatusBadRequest)
		return
	}

	fmt.Println(emp)

	if _, exists := r.employees[emp.ID()]; exists {
		fmt.Println("Employee Already Exists")
		http.Error(w, "Invalid Request body", http.StatusUnprocessableEntity)
		return
	}

	r.employees[emp.ID()] = emp

	r.ListAllEmployees()

	w.WriteHeader(http.StatusCreated)

}

func (r *EmployeeRepo) ReadEmployee(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	idParam := query.Get("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	emp, exists := r.employees[id]
	if !exists {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emp)
}

func (r *EmployeeRepo) UpdateEmployee(w http.ResponseWriter, req *http.Request) {
	var emp models.Employee

	if err := json.NewDecoder(req.Body).Decode(&emp); err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	if _, exists := r.employees[emp.ID()]; !exists {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	r.employees[emp.ID()] = emp
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Employee Updated: %+v", emp)
}

func (r *EmployeeRepo) DeleteEmployee(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	idParam := query.Get("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID Parameter", http.StatusBadRequest)
		return
	}

	if _, exists := r.employees[id]; !exists {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	delete(r.employees, id)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Employee with ID %d deleted successfully", id)
}

/*
func (r *EmployeeRepo) GetEmployee(id int) (models.Employee, error) {

	if val, exists := r.employees[id]; !exists {
		fmt.Println("No Employee Found")
		return models.Employee{}, errors.New("Invalid Employee ID")
	} else {
		r.ListAllEmployees()
		return val, nil
	}

}

func (r *EmployeeRepo) UpdateEmployee(e models.Employee) error {

	if _, exists := r.employees[e.ID()]; !exists {
		fmt.Println("No Employee Found to update")
		return errors.New("Invalid Employee ID")
	}

	r.employees[e.ID()] = e

	r.ListAllEmployees()
	return nil
}

func (r *EmployeeRepo) DeleteEmployee(id int) error {

	if _, exists := r.employees[id]; !exists {
		fmt.Println("No Employee found to delete")
		return errors.New("Not a new Employee")
	}

	delete(r.employees, id)

	r.ListAllEmployees()
	return nil
}
*/

func (r *EmployeeRepo) ListAllEmployees() {
	fmt.Println(r.employees)
}
