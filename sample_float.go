package main

import "fmt"

func main() {
	n := 3.14
	fmt.Printf("%v, %T\n", n, n)
	n = 13.7e72
	fmt.Printf("%v, %T\n", n, n)
	n = 2.1e14
	fmt.Printf("%v, %T\n", n, n)

	fmt.Println("................")
	a := 10.2
	b := 3.7
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)

	fmt.Println("................")

	var c1 complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", c1, c1)
	fmt.Printf("%v, %T\n", real(c1), real(c1))
	fmt.Printf("%v, %T\n", imag(c1), imag(c1))

	fmt.Println("................")

	c1 = 2i
	fmt.Printf("%v, %T\n", c1, c1)

	c2 := 2 + 5.2i

	c3 := complex128(c1) + c2
	fmt.Printf("%v, %T\n", c3, c3)

	fmt.Println("................")
	var c4 complex128 = complex(5, 12)
	fmt.Printf("%v, %T\n", c4, c4)
}
