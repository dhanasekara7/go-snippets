package main

import "fmt"

func main() {
	fmt.Println("start")
	defer fmt.Println("deferred then ")
	panic("panic !!!!")
	fmt.Println("end")
}
