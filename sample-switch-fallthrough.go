package main

import "fmt"

func main() {

	i := 5
	switch {
	case i <= 10:
		fmt.Println("i<=10")
		fallthrough // bypass case check
	case i < 10:
		fmt.Println("i<10")
		fallthrough
	default:
		fmt.Println("default ...")
	}
}
