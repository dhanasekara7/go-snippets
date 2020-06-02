package main

import (
	"fmt"
	"runtime"
)

func main() {
	//runtime.GOMAXPROCS(100)
	fmt.Println(runtime.GOMAXPROCS(-1))
}
