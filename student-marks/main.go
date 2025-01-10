package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Marks []int
}

func (s *Student) AddMark(mark int) {
	s.Marks = append(s.Marks, mark)
}

func (s Student) CalculateAverage() float64 {
	if len(s.Marks) == 0 {
		return 0.0
	}

	sum := 0
	for _, mark := range s.Marks {
		sum += mark
	}

	return float64(sum) / float64(len(s.Marks))
}

func main() {
	s := Student{Name: "Ram"}

	s.AddMark(63)
	s.AddMark(72)
	s.AddMark(81)

	average := s.CalculateAverage()

	fmt.Printf("The average marks for %s are: %.2f\n", s.Name, average)
}
