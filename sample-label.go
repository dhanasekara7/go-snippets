package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 2 {
				break
			}
			fmt.Println("i=", i, "j=", j)
		}
	}
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>")
Loop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 2 {
				break Loop
			}
			fmt.Println("i=", i, "j=", j)
		}
	}
}
