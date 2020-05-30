package main

import "fmt"

func main() {
	findSum(1, 2, 3, 4, 5)
}

// int return type
func findSum(values ...int) (result int) {
	fmt.Println(values)
	result = 0
	for k, v := range values {
		fmt.Println("K=", k)
		result += v

	}
	fmt.Println("The sum is=", result)
}
