package main

import "fmt"

func main() {
	i := 0

	for i < 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println(">>>>>>>>>>>")

	for {
		if i > 5 {
			break
		}
		fmt.Println(i)
		i++
	}
}
