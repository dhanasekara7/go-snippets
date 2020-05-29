package main

import "fmt"

func main() {

	i := 10
	switch {
	case i <= 10:
		fmt.Println("i<=10")
	case i > 10:
		fmt.Println("i>10")
	default:
		fmt.Println("default ...")
	}
}
