package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// can not use 2 inits in for loop
	//for i:= 0, j:=0; i < 10; i++,j++ {
	//	fmt.Println(i)
	//}

	for i, j := 0, 0; i < 10; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}
}
