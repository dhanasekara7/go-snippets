package main

import "fmt"

func main() {
	a := 10
	b := 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	fmt.Println("..............")
	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	fmt.Println(a &^ b)
	fmt.Println("..............")

	var c int = 10
	var d int8 = 3
	//fmt.Println(c + d)
	// mismatched types error, because int + int8
	fmt.Println(c + int(d))

	s := 8
	fmt.Println(s << 3)
	fmt.Println(s >> 3)
	fmt.Println(s)

}
