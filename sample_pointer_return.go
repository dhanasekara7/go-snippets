package main

import "fmt"

func main() {
	xs := sumMe(1, 2, 3, 4, 5)
	fmt.Println(*xs)
}

func sumMe(integers ...int) *int {
	result := 0
	for _, v := range integers {
		result += v
	}
	return &result
}
