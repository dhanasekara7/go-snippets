package main

import "fmt"

func main() {
	s := "this is a string"
	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b)

	// single quote
	r := 'a' // rune
	fmt.Printf("%v, %T\n", r, r)

	// same result
	var r2 rune = 'b'
	fmt.Printf("%v, %T\n", r2, r2)

}
