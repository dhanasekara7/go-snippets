package main

import "fmt"

type myStruct struct {
	foo int
}

func main() {
	var ms *myStruct
	ms = &myStruct{foo: 42}
	fmt.Println(ms)
}
