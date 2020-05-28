package main

import (
	"fmt"
)

// const shandowing
const shConst int16 = 12

func main() {

	const myConst int = 4
	fmt.Printf("%v, %T\n", myConst, myConst)

	//cant set the following
	//const piConst float64 = math.Sin(1.57)

	const a int = 14
	const b string = "foo"
	const c float32 = 3.14
	const d bool = true
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)

	fmt.Printf("%v, %T\n", shConst, shConst)
	const shConst int = 12
	fmt.Printf("%v, %T\n", shConst, shConst)

	const inferType = 10 // infer as int
	fmt.Printf("%v, %T\n", inferType, inferType)

	const inferTypeAsInt16 = 10 // infer as int, but during addition to other var, infer relevant data type
	var x int16 = 20
	fmt.Printf("%v, %T\n", inferTypeAsInt16, inferTypeAsInt16)
	fmt.Printf("%v, %T\n", x+inferTypeAsInt16, x+inferTypeAsInt16)

}
