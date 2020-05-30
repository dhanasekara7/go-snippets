package main

import "fmt"

func main() {
	defer fmt.Println("this")
	defer fmt.Println("is")
	defer fmt.Println("fun")
}
