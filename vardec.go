package main

import "fmt"

var actorname string = "funactor"
var i int
var ii int = 123

//i = 42

var (
	a string = "alphabets"
	b int    = 120
)

func main() {
	i = 42
	var j int = 53
	k := 64
	var l float32 = 23
	m := 74.

	fmt.Println("this is var dec=", i, j, k, l)
	fmt.Printf("%v, %T\n", k, k)
	fmt.Printf("%v, %T\n", m, m)
	fmt.Printf("%v, %T\n", actorname, actorname)
}
