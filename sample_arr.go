package main

import "fmt"

func main() {
	grades := [3]int{97, 32, 42}
	fmt.Printf("%v\n", grades)

	betterGrades := [...]int{97, 32, 42}
	fmt.Printf("%v\n", betterGrades)

	bestGrades := []int{97, 32, 42}
	fmt.Printf("%v\n", bestGrades)

	var students [3]string
	fmt.Printf("%v\n", students)

	var bestStudents [3]string
	bestStudents[0] = "dhans"
	bestStudents[1] = "sekar"
	bestStudents[2] = "vivek"
	fmt.Printf("%v\n", bestStudents)
	fmt.Printf("%v\n", len(bestStudents))

}
