package main

import "fmt"

func main() {
	var a int = 2
	var b *int = &a
	fmt.Println(a, *b)
	a = 3
	fmt.Println(a, *b)
	*b = 23
	fmt.Println(a, *b)

}
