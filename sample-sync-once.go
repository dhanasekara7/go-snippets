package main

import (
	"fmt"
	"sync"
)

var on sync.Once

func Xinit() {
	fmt.Println("this is init")
}

func main() {
	on.Do(Xinit)
	fmt.Println("this is main")
}
