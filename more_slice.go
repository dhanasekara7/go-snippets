package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:] // refers reference, but change this to a, will not refer reference, depending on a := [] or [...] or [10]
	c := a[3:]
	d := a[:6]
	e := a[3:6]
	b[5] = 42
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

}
