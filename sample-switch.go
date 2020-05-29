package main

import "fmt"

func main() {
	var i int = 4
	//switch i {
	switch i = i - 2; i {
	case 1:
		fmt.Println(">>>1")
	case 2, 4:
		fmt.Println(">>>", i)
	case 3:
		fmt.Println(">>>2")
	default:
		fmt.Println(">>>3")
	}
	fmt.Println("i=", i)
}
