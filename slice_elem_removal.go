package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}

	b := a[1:]
	fmt.Printf("%v\n", b)
	fmt.Printf("a=%v\n", a) // oops

	b = a[:len(a)-1]
	fmt.Printf("%v\n", b)
	fmt.Printf("a=%v\n", a) // oops

	b = append(a[:2], a[3:]...) // notice the ...
	fmt.Printf("%v\n", b)

	fmt.Printf("a=%v\n", a) // oops// oops

}
