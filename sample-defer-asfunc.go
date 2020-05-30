package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error11", err)
		}
	}()

	panic("panic !!!!")
	fmt.Println("end")
}
