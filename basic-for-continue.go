package main

import (
	"fmt"
)

func main() {
	i := 0
	for {
		i++
		fmt.Println(">>>>i=", i)
		if i < 5 {
			continue
		} else {
			break
		}

	}
}
