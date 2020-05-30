package main

import "fmt"

func main() {
	sums := sumMe(1, 2, 3, 4, 5)
	fmt.Println(sums)
}

func sumMe(integers ...int) (result int) {
	for _, v := range integers {
		result += v
	}
	return
}
