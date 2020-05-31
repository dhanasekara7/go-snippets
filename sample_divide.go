package main

import "fmt"

func main() {
	result := divide(1.0, 2.0)
	fmt.Println(result)
}

func divide(a, b float64) (result float64) {
	if b == 0 {
		panic("can not divide by zero")
	}
	result = a / b
	return result
}
