package main

import "fmt"

type Student struct {
	Name  string
	Marks []int
}

func AddMark(student *Student, mark int) {
	student.Marks = append(student.Marks, mark)
}

func CalculateAverage(student Student) float64 {
	if len(student.Marks) == 0 {
		return 0.0
	}

	sum := 0
	for index, mark := range student.Marks { //using "index" instead of "_" (just out of respect towards the indexes)
		sum += mark
		fmt.Scanln(index)
	}

	return float64(sum) / float64(len(student.Marks))
}

func main() {
	student := Student{Name: "Rama"}

	AddMark(&student, 63)
	AddMark(&student, 72)
	AddMark(&student, 81)

	average := CalculateAverage(student)

	fmt.Printf("The average marks for %s are: %.2f\n", student.Name, average)
}
