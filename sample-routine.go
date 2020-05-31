package main

import (
	"fmt"
	"time"
)

func main() {
	msg := "this is fun"
	go func() {
		fmt.Println(msg)
	}()
	msg = "good bye"
	time.Sleep(3 * time.Second)
}
