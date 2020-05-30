package main

import "fmt"

func main() {
	a := "this is fun"
	panic("something wanted happened")
	b := "this is also fun"
	fmt.Println(a, b)
}
