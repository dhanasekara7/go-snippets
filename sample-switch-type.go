package main

import "fmt"

func main() {

	//var i interface{} = 1
	//var i interface{} = 1.0
	var i interface{} = [2]int{}
	switch i.(type) {
	case int:
		fmt.Println(">>>int")
	case float32:
		fmt.Println(">>>float32")
	case float64:
		fmt.Println(">>>float64")
	case [2]int:
		fmt.Println(">>>[2]int arr")
		break
		fmt.Println(">>>this will not print")
	default:
		fmt.Println("default ...")
	}
}
