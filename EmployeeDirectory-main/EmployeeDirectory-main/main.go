package main

import (
	"employeeeDirectory/repository"
	"fmt"
	"log"
	"net/http"
)

func main() {
	repo := repository.NewEmployeeRepo()

	http.HandleFunc("/create", repo.CreateEmployee)
	http.HandleFunc("/read", repo.ReadEmployee)
	http.HandleFunc("/update", repo.UpdateEmployee)
	http.HandleFunc("/delete", repo.DeleteEmployee)

	fmt.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

/*
CRUD

*/

/*
func main() {

	repo := repository.NewEmployeeRepo()

	Execute(repo)

}

func Execute(repo service.EmployeeService) {

	//Create

	http.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodPost:
			{

				repo.CreateEmployee(w, r)

			}
		default:
			{
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}

		}
	})

	fmt.Println("Starting Server")


	http.ListenAndServe(":8080", nil)
*/

/*
		//Update

		err := repo.UpdateEmployee(models.Employee{
			EmployeeID:   2,
			EmployeeName: "Bhavani",
			EmployeeAge:  30,
		})

		if err != nil {
			repo.ListAllEmployees()
			fmt.Println(err)
		} else {
			fmt.Println("Employee Updated Successfully ")
		}

		//GET
		fmt.Println("**********Getting an Employee with ID 2***************")
		val, err := repo.GetEmployee(2)

		fmt.Println(val)

		if err != nil {
			repo.ListAllEmployees()
			fmt.Println(err)
		} else {
			fmt.Println("Employee Updated Successfully ")
		}

		//Delete

		fmt.Println("**********Deleting an Employee with ID 2***************")
		err = repo.DeleteEmployee(2)

		if err != nil {
			repo.ListAllEmployees()
			fmt.Println(err)
		} else {
			fmt.Println("Employee Updated Successfully ")
		}


}
*/
