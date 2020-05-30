package main

import "fmt"

func main() {
	// defer will take the argument during defer call

	a := "Start"
	defer fmt.Println(a)
	a = "end"
}
