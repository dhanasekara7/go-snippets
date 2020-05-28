package main

import "fmt"

func main() {
	fmt.Println(">>>>>array with ... size")
	a1 := [...]int{1, 2, 3}
	b1 := a1

	b1[1] = 5

	fmt.Printf("%v\n", a1)
	fmt.Printf("%v\n", b1)

	fmt.Println(">>>>>array with defined size")
	a2 := [3]int{1, 2, 3}
	b2 := a2

	b2[1] = 5

	fmt.Printf("%v\n", a2)
	fmt.Printf("%v\n", b2)

	fmt.Println(">>>>>array(undefined size) reference")
	// array reference
	aa := []int{1, 2, 3}
	bb := aa

	bb[1] = 5

	fmt.Printf("%v\n", aa)
	fmt.Printf("%v\n", bb)

	fmt.Println(">>>>>array referencing using &")
	a3 := [...]int{1, 2, 3}
	b3 := &a3

	b3[1] = 5

	fmt.Printf("%v\n", a3)
	fmt.Printf("%v\n", *b3)
}
