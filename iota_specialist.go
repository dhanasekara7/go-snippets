package main

import "fmt"

// _ = iota
// also works
const (
	error = iota
	// error = iota + 5 // for offset
	dogSpecialist
	catSpecialist
	snakeSpecialist
)

func main() {
	var splType int
	fmt.Printf("%v\n", splType == dogSpecialist)
	// if error const is not defined the above statement will be true
	// because splType assigment forgotten
}
