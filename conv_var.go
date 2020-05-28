package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	var j float32
	j = float32(i)
	fmt.Printf("%v, %T\n", j, j)

	var k float32 = 32.5
	fmt.Printf("%v, %T\n", k, k)

	// losing value
	var l int = int(k)
	fmt.Printf("%v, %T\n", l, l)

	var m int = 42
	var s string = string(m)
	fmt.Printf("%v, %T\n", s, s)

	var ss string = strconv.Itoa(m)
	fmt.Printf("%v, %T\n", ss, ss)

}
