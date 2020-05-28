package main

import "fmt"

func main() {
	var x bool
	fmt.Printf("%v, %T\n", x, x)

	m := 1 == 2
	n := 1 == 1
	fmt.Printf("%v, %T\n", m, m)
	fmt.Printf("%v, %T\n", n, n)

	var un uint = 42
	fmt.Printf("%v, %T\n", un, un)

}
