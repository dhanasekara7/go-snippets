package main

import "fmt"

func main() {
	paramx := "this"
	paramy := "fun"
	method1(paramx, paramy)
}

func method1(x, y string) {
	fmt.Println("x=", x)
	fmt.Println("y=", y)

}
