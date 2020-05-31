package main

import "fmt"

func main() {
	inc := Incrementor{}
	res := inc.inc(10)
	fmt.Println(res)
}

type increment interface {
	inc(x int) int
}

type Incrementor struct {
}

func (incrementor Incrementor) inc(x int) int {
	x = x + 1
	return x
}
