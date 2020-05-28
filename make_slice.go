package main

import "fmt"

func main() {
	a := make([]int, 3)
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	b := make([]int, 3, 100)
	fmt.Println(b)
	fmt.Printf("Length: %v\n", len(b))
	fmt.Printf("Capacity: %v\n", cap(b))
}
