package main

import "fmt"

const (
	aa = iota
	bb = iota
	cc = iota
)

// iota reset
const (
	aaa = iota
	bbb = iota
	ccc = iota
)

//iota from prev line
const (
	aaaa = iota
	bbbb
	cccc
)

func main() {
	fmt.Printf("%v \n", aa)
	fmt.Printf("%v \n", bb)
	fmt.Printf("%v \n", cc)

	fmt.Println(".............")

	fmt.Printf("%v \n", aaa)
	fmt.Printf("%v \n", bbb)
	fmt.Printf("%v \n", ccc)

	fmt.Println(".............")

	fmt.Printf("%v \n", aaaa)
	fmt.Printf("%v \n", bbbb)
	fmt.Printf("%v \n", cccc)

}
