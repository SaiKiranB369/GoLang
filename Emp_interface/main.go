package main

import "fmt"

//Employee struct to group the employee details
type Employee struct {
	Name        string
	Designation string
	ID          int
	Salary      float64
}

//EmployeeManger interface to manage the CRUD operations
type EmployeeManager interface {
	Create(employee Employee)
	Read(id int) *Employee
	Update(id int, updatedEmployee Employee) bool
	Delete(id int) bool
	ListAll()
}

//EmployeeDB struct to store the employees
type EmployeeDB struct {
	employees []Employee
}

//Create Employee
func (db *EmployeeDB) Create(employee Employee) {
	db.employees = append(db.employees, employee)
	fmt.Println("Employee added successfully.")
}

//Read Employee
func (db *EmployeeDB) Read(id int) *Employee {
	for _, emp := range db.employees {
		if emp.ID == id {
			return &emp
		}
	}
	return nil
}

//Update Employee
func (db *EmployeeDB) Update(id int, updatedEmployee Employee) bool {
	for i, emp := range db.employees {
		if emp.ID == id {
			db.employees[i] = updatedEmployee
			fmt.Println("Employee updated successfully.")
			return true
		}
	}
	return false
}

//Delete Employee
func (db *EmployeeDB) Delete(id int) bool {
	for i, emp := range db.employees {
		if emp.ID == id {
			db.employees = append(db.employees[:i], db.employees[i+1:]...)
			fmt.Println("Employee deleted successfully.")
			return true
		}
	}
	return false
}

// ListAll prints all employees
func (db *EmployeeDB) ListAll() {
	fmt.Println("Employee List:")
	for _, emp := range db.employees {
		fmt.Printf("ID: %d, Name: %s, Designation: %s, Salary: %.2f\n", emp.ID, emp.Name, emp.Designation, emp.Salary)
	}
}

func main() {
	var manager EmployeeManager = &EmployeeDB{}

	// Create employees
	manager.Create(Employee{Name: "Rama", Designation: "Software Engineer", ID: 1, Salary: 140000})
	manager.Create(Employee{Name: "Vasudeva", Designation: "Data Scientist", ID: 2, Salary: 80000})

	// List all employees
	manager.ListAll()

	// Read an employee by ID
	id := 1
	emp := manager.Read(id)
	if emp != nil {
		fmt.Printf("Details of employee with ID %d: %+v\n", id, *emp)
	} else {
		fmt.Printf("Employee with ID %d not found.\n", id)
	}

	// Update an employee
	updated := manager.Update(1, Employee{Name: "John Doe", Designation: "Senior Software Engineer", ID: 1, Salary: 90000})
	if !updated {
		fmt.Printf("Failed to update employee with ID %d.\n", id)
	}

	// List all employees after update
	manager.ListAll()

	// Delete an employee
	deleted := manager.Delete(2)
	if !deleted {
		fmt.Printf("Failed to delete employee with ID %d.\n", 2)
	}

	// List all employees after deletion
	manager.ListAll()
}
