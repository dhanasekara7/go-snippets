package main

import "fmt"

func main() {
	for i := 0; i < 15; i++ {
		func(i int) {
			fmt.Println("i=", i)
		}(i)
	}
}
