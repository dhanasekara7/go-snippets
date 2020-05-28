package main

import "fmt"

func main() {
	var i int
	i = 42
	var j int = 53
	k := 64
	var l float32 = 23
	fmt.Println("this is var dec=", i, j, k, l)
	fmt.Printf("%v %T", k, k)
}
